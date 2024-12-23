package answer

import (
	"context"
	"github.com/upassed/upassed-answer-service/internal/config"
	domain "github.com/upassed/upassed-answer-service/internal/repository/model"
	business "github.com/upassed/upassed-answer-service/internal/service/model"
	"log/slog"
)

type Service interface {
	Create(ctx context.Context, answer *business.Answer) (*business.AnswerCreateResponse, error)
}

type serviceImpl struct {
	cfg        *config.Config
	log        *slog.Logger
	repository repository
}

type repository interface {
	Save(ctx context.Context, answers []*domain.Answer) error
}

func New(cfg *config.Config, log *slog.Logger, repository repository) Service {
	return &serviceImpl{
		cfg:        cfg,
		log:        log,
		repository: repository,
	}
}
