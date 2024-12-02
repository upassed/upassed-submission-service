package answer

import (
	"github.com/upassed/upassed-answer-service/internal/config"
	"github.com/upassed/upassed-answer-service/internal/logging"
	"github.com/upassed/upassed-answer-service/internal/middleware/common/auth"
	"github.com/wagslane/go-rabbitmq"
	"log/slog"
)

type rabbitClient struct {
	authClient       auth.Client
	rabbitConnection *rabbitmq.Conn
	cfg              *config.Config
	log              *slog.Logger
}

func Initialize(authClient auth.Client, rabbitConnection *rabbitmq.Conn, cfg *config.Config, log *slog.Logger) {
	log = logging.Wrap(log,
		logging.WithOp(Initialize),
	)

	client := &rabbitClient{
		authClient:       authClient,
		rabbitConnection: rabbitConnection,
		cfg:              cfg,
		log:              log,
	}

	go func() {
		if err := InitializeCreateQueueConsumer(client); err != nil {
			log.Error("error while initializing answer queue consumer", logging.Error(err))
			return
		}
	}()
}
