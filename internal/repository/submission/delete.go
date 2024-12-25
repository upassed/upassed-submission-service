package submission

import (
	"context"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
)

func (repository *repositoryImpl) Delete(_ context.Context, _ *domain.SubmissionDeleteParams) error {
	panic("implement me")
}
