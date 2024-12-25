package submission

import (
	"context"
	"github.com/upassed/upassed-submission-service/internal/config"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
	"log/slog"
)

type Service interface {
	Create(ctx context.Context, submission *business.Submission) (*business.SubmissionCreateResponse, error)
}

type serviceImpl struct {
	cfg        *config.Config
	log        *slog.Logger
	repository repository
}

type repository interface {
	Save(ctx context.Context, submissions []*domain.Submission) error
	Exists(ctx context.Context, params *domain.SubmissionExistCheckParams) (bool, error)
	Delete(ctx context.Context, params *domain.SubmissionDeleteParams) error
}

func New(cfg *config.Config, log *slog.Logger, repository repository) Service {
	return &serviceImpl{
		cfg:        cfg,
		log:        log,
		repository: repository,
	}
}
