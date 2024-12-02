package app

import (
	"github.com/upassed/upassed-answer-service/internal/config"
	"github.com/upassed/upassed-answer-service/internal/logging"
	"github.com/upassed/upassed-answer-service/internal/messanging"
	answerRabbit "github.com/upassed/upassed-answer-service/internal/messanging/answer"
	"github.com/upassed/upassed-answer-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-answer-service/internal/repository"
	"github.com/upassed/upassed-answer-service/internal/server"
	"log/slog"
)

type App struct {
	Server *server.AppServer
}

func New(config *config.Config, log *slog.Logger) (*App, error) {
	log = logging.Wrap(log, logging.WithOp(New))

	_, err := repository.OpenGormDbConnection(config, log)
	if err != nil {
		return nil, err
	}

	rabbit, err := messanging.OpenRabbitConnection(config, log)
	if err != nil {
		return nil, err
	}

	authClient, err := auth.NewClient(config, log)
	if err != nil {
		return nil, err
	}

	answerRabbit.Initialize(authClient, rabbit, config, log)
	appServer := server.New(server.AppServerCreateParams{
		Config:     config,
		Log:        log,
		AuthClient: authClient,
	})

	log.Info("app successfully created")
	return &App{
		Server: appServer,
	}, nil
}
