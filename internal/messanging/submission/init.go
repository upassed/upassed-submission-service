package submission

import (
	"errors"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/wagslane/go-rabbitmq"
)

var (
	errCreatingSubmissionCreateQueueConsumer = errors.New("unable to create submission queue consumer")
	errRunningSubmissionCreateQueueConsumer  = errors.New("unable to run submission queue consumer")
)

func InitializeCreateQueueConsumer(client *rabbitClient) error {
	log := logging.Wrap(client.log,
		logging.WithOp(InitializeCreateQueueConsumer),
	)

	log.Info("started crating submission create queue consumer")
	submissionCreateGroupConsumer, err := rabbitmq.NewConsumer(
		client.rabbitConnection,
		client.cfg.Rabbit.Queues.SubmissionCreate.Name,
		rabbitmq.WithConsumerOptionsRoutingKey(client.cfg.Rabbit.Queues.SubmissionCreate.RoutingKey),
		rabbitmq.WithConsumerOptionsExchangeName(client.cfg.Rabbit.Exchange.Name),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)

	if err != nil {
		log.Error("unable to create submission queue consumer", logging.Error(err))
		return errCreatingSubmissionCreateQueueConsumer
	}

	defer submissionCreateGroupConsumer.Close()
	if err := submissionCreateGroupConsumer.Run(client.CreateQueueConsumer()); err != nil {
		log.Error("unable to run submission queue consumer")
		return errRunningSubmissionCreateQueueConsumer
	}

	log.Info("submission queue consumer successfully initialized")
	return nil
}
