package submission

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-submission-service/internal/config"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
	"github.com/upassed/upassed-submission-service/pkg/client"
	"google.golang.org/grpc"
)

type submissionServerAPI struct {
	client.UnimplementedSubmissionServer
	cfg     *config.Config
	service assignmentService
}

type assignmentService interface {
	FindByFormID(ctx context.Context, formID uuid.UUID) (*business.FormSubmissions, error)
}

func Register(gRPC *grpc.Server, cfg *config.Config, service assignmentService) {
	client.RegisterSubmissionServer(gRPC, &submissionServerAPI{
		cfg:     cfg,
		service: service,
	})
}
