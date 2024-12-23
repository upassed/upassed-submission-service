package submission

import (
	"context"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/amqp"
	loggingMiddleware "github.com/upassed/upassed-submission-service/internal/middleware/amqp/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/amqp/recovery"
	requestidMiddleware "github.com/upassed/upassed-submission-service/internal/middleware/amqp/request_id"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	requestid "github.com/upassed/upassed-submission-service/internal/middleware/common/request_id"
	"github.com/upassed/upassed-submission-service/internal/tracing"
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

		log.Info("consumed submission create message", slog.String("messageBody", string(delivery.Body)))
		spanContext, span := otel.Tracer(client.cfg.Tracing.SubmissionTracerName).Start(ctx, "submission#Create")
		span.SetAttributes(attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)))
		defer span.End()

		log.Info("converting message body to submission create request struct")
		request, err := ConvertToSubmissionCreateRequest(delivery.Body)
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

		log.Info("validating submission create request")
		if err := request.Validate(); err != nil {
			log.Error("submission create request is invalid", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		log.Info("creating submission")
		response, err := client.service.Create(spanContext, ConvertToBusinessSubmission(request, studentUsername))
		if err != nil {
			log.Error("unable to create submission", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		log.Info("successfully created submission", slog.Any("createdSubmissionIDs", response.CreatedSubmissionIDs))
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
