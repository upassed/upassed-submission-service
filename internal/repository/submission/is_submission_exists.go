package submission

import (
	"context"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
)

func (repository *repositoryImpl) Exists(_ context.Context, _ *domain.SubmissionExistCheckParams) (bool, error) {
	panic("implement me")
}
