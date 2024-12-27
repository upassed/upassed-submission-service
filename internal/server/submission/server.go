package submission

import (
	"context"
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
	FindStudentFormSubmissions(ctx context.Context, params *business.StudentFormSubmissionSearchParams) (*business.FormSubmissions, error)
}

func Register(gRPC *grpc.Server, cfg *config.Config, service assignmentService) {
	client.RegisterSubmissionServer(gRPC, &submissionServerAPI{
		cfg:     cfg,
		service: service,
	})
}
