package answer

import (
	"context"
	"github.com/upassed/upassed-answer-service/internal/logging"
	"github.com/upassed/upassed-answer-service/internal/middleware/amqp"
	loggingMiddleware "github.com/upassed/upassed-answer-service/internal/middleware/amqp/logging"
	"github.com/upassed/upassed-answer-service/internal/middleware/amqp/recovery"
	requestidMiddleware "github.com/upassed/upassed-answer-service/internal/middleware/amqp/request_id"
	"github.com/upassed/upassed-answer-service/internal/middleware/common/auth"
	requestid "github.com/upassed/upassed-answer-service/internal/middleware/common/request_id"
	"github.com/upassed/upassed-answer-service/internal/tracing"
	"github.com/wagslane/go-rabbitmq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"log/slog"
)

func (client *rabbitClient) CreateQueueConsumer() rabbitmq.Handler {
	baseHandler := func(ctx context.Context, delivery rabbitmq.Delivery) rabbitmq.Action {
		log := logging.Wrap(client.log,
			logging.WithOp(client.CreateQueueConsumer),
			logging.WithCtx(ctx),
		)

		log.Info("consumed answers create message", slog.String("messageBody", string(delivery.Body)))
		spanContext, span := otel.Tracer(client.cfg.Tracing.AnswerTracerName).Start(ctx, "answer#Create")
		span.SetAttributes(attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)))
		defer span.End()

		log.Info("converting message body to answer create request struct")
		request, err := ConvertToAnswerCreateRequest(delivery.Body)
		if err != nil {
			log.Error("unable to convert message body to create request struct", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		studentUsername := ctx.Value(auth.UsernameKey).(string)
		span.SetAttributes(
			attribute.String("formID", request.FormID),
			attribute.String("questionID", request.QuestionID),
			attribute.String("studentUsername", studentUsername),
		)

		log.Info("validating answer create request")
		if err := request.Validate(); err != nil {
			log.Error("answer create request is invalid", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		log.Info("creating answer")
		response, err := client.service.Create(spanContext, ConvertToBusinessAnswer(request, studentUsername))
		if err != nil {
			log.Error("unable to create answer", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		log.Info("successfully created answer", slog.Any("createdAnswerID", response.CreatedAnswerIDs))
		return rabbitmq.Ack
	}

	handlerWithMiddleware := amqp.ChainMiddleware(
		baseHandler,
		requestidMiddleware.Middleware(),
		loggingMiddleware.Middleware(client.log),
		recovery.Middleware(client.log),
		client.authClient.AmqpMiddleware(client.cfg, client.log),
	)

	return func(d rabbitmq.Delivery) (action rabbitmq.Action) {
		ctx := context.Background()
		return handlerWithMiddleware(ctx, d)
	}
}
