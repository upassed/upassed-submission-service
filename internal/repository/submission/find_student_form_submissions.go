package submission

import (
	"context"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
)

func (repository *repositoryImpl) FindStudentFormSubmissions(_ context.Context, _ *domain.StudentFormSubmissionsSearchParams) ([]*domain.Submission, error) {
	//TODO implement me
	panic("implement me")
}
