package submission

import (
	"context"
	"github.com/google/uuid"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
)

func (service *serviceImpl) FindByFormID(_ context.Context, _ uuid.UUID) (*business.FormSubmissions, error) {
	//TODO implement me
	panic("implement me")
}
