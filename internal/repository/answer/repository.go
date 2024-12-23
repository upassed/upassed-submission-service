package answer

import (
	"context"
	"github.com/upassed/upassed-answer-service/internal/config"
	domain "github.com/upassed/upassed-answer-service/internal/repository/model"
	"gorm.io/gorm"
	"log/slog"
)

type Repository interface {
	Save(ctx context.Context, answers []*domain.Answer) error
}

type repositoryImpl struct {
	db  *gorm.DB
	cfg *config.Config
	log *slog.Logger
}

func New(db *gorm.DB, cfg *config.Config, log *slog.Logger) Repository {
	return &repositoryImpl{
		db:  db,
		cfg: cfg,
		log: log,
	}
}
