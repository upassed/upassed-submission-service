package answer

import (
	"errors"
	"github.com/upassed/upassed-answer-service/internal/logging"
	"github.com/wagslane/go-rabbitmq"
)

var (
	errCreatingAnswerCreateQueueConsumer = errors.New("unable to create answer queue consumer")
	errRunningAnswerCreateQueueConsumer  = errors.New("unable to run answer queue consumer")
)

func InitializeCreateQueueConsumer(client *rabbitClient) error {
	log := logging.Wrap(client.log,
		logging.WithOp(InitializeCreateQueueConsumer),
	)

	log.Info("started crating answer create queue consumer")
	answerCreateGroupConsumer, err := rabbitmq.NewConsumer(
		client.rabbitConnection,
		client.cfg.Rabbit.Queues.AnswerCreate.Name,
		rabbitmq.WithConsumerOptionsRoutingKey(client.cfg.Rabbit.Queues.AnswerCreate.RoutingKey),
		rabbitmq.WithConsumerOptionsExchangeName(client.cfg.Rabbit.Exchange.Name),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)

	if err != nil {
		log.Error("unable to create answer queue consumer", logging.Error(err))
		return errCreatingAnswerCreateQueueConsumer
	}

	defer answerCreateGroupConsumer.Close()
	if err := answerCreateGroupConsumer.Run(client.CreateQueueConsumer()); err != nil {
		log.Error("unable to run answer queue consumer")
		return errRunningAnswerCreateQueueConsumer
	}

	log.Info("answer queue consumer successfully initialized")
	return nil
}
