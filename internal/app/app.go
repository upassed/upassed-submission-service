package app

import (
	"github.com/upassed/upassed-submission-service/internal/config"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/messanging"
	submissionRabbit "github.com/upassed/upassed-submission-service/internal/messanging/submission"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-submission-service/internal/repository"
	submissionRepository "github.com/upassed/upassed-submission-service/internal/repository/submission"
	"github.com/upassed/upassed-submission-service/internal/server"
	"github.com/upassed/upassed-submission-service/internal/service/submission"
	"log/slog"
)

type App struct {
	Server *server.AppServer
}

func New(config *config.Config, log *slog.Logger) (*App, error) {
	log = logging.Wrap(log, logging.WithOp(New))

	db, err := repository.OpenGormDbConnection(config, log)
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

	submissionRepo := submissionRepository.New(db, config, log)
	submissionService := submission.New(config, log, submissionRepo)

	submissionRabbit.Initialize(authClient, submissionService, rabbit, config, log)
	appServer := server.New(server.AppServerCreateParams{
		Config:            config,
		Log:               log,
		AuthClient:        authClient,
		SubmissionService: submissionService,
	})

	log.Info("app successfully created")
	return &App{
		Server: appServer,
	}, nil
}
