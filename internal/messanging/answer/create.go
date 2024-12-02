package answer

import (
	"context"
	"github.com/upassed/upassed-answer-service/internal/logging"
	"github.com/upassed/upassed-answer-service/internal/middleware/amqp"
	loggingMiddleware "github.com/upassed/upassed-answer-service/internal/middleware/amqp/logging"
	"github.com/upassed/upassed-answer-service/internal/middleware/amqp/recovery"
	requestidMiddleware "github.com/upassed/upassed-answer-service/internal/middleware/amqp/request_id"
	requestid "github.com/upassed/upassed-answer-service/internal/middleware/common/request_id"
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

		log.Info("consumed answer create message", slog.String("messageBody", string(delivery.Body)))
		_, span := otel.Tracer(client.cfg.Tracing.AnswerTracerName).Start(ctx, "answer#Create")
		span.SetAttributes(attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)))
		defer span.End()

		log.Info("converting message body to answer create request struct")
		log.Info("successfully created answer", slog.Any("createdAnswerID", 1))
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
