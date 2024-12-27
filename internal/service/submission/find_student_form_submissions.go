package submission

import (
	"context"
	"errors"
	"github.com/upassed/upassed-submission-service/internal/async"
	"github.com/upassed/upassed-submission-service/internal/handling"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
	"github.com/upassed/upassed-submission-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

var (
	ErrSubmissionSearchingDeadlineExceeded = errors.New("deadline exceeded while searching student's form submissions")
)

func (service *serviceImpl) FindStudentFormSubmissions(ctx context.Context, params *business.StudentFormSubmissionSearchParams) (*business.FormSubmissions, error) {
	username := ctx.Value(auth.UsernameKey).(string)

	spanContext, span := otel.Tracer(service.cfg.Tracing.SubmissionTracerName).Start(ctx, "submissionService#FindStudentFormSubmissions")
	span.SetAttributes(
		attribute.String("formID", params.FormID.String()),
		attribute.String(auth.UsernameKey, username),
	)
	defer span.End()

	log := logging.Wrap(service.log,
		logging.WithOp(service.FindStudentFormSubmissions),
		logging.WithCtx(ctx),
		logging.WithAny("formID", params.FormID.String()),
	)

	log.Info("started searching student's form submissions")
	timeout := service.cfg.GetEndpointExecutionTimeout()

	foundFormSubmissions, err := async.ExecuteWithTimeout(spanContext, timeout, func(ctx context.Context) (*business.FormSubmissions, error) {
		log.Info("searching submissions data in the database")
		searchParams := ConvertToStudentFormSubmissionsSearchParams(params.StudentUsername, params.FormID)
		foundSubmissions, err := service.repository.FindStudentFormSubmissions(ctx, searchParams)

		if err != nil {
			log.Error("unable to find student's form submissions in the database", logging.Error(err))
			tracing.SetSpanError(span, err)
			return nil, err
		}

		return ConvertToFormSubmissions(foundSubmissions), nil
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("submissions searching deadline exceeded")
			tracing.SetSpanError(span, err)
			return nil, handling.Wrap(ErrSubmissionSearchingDeadlineExceeded, handling.WithCode(codes.DeadlineExceeded))
		}

		log.Error("error while searching student's form submissions", logging.Error(err))
		tracing.SetSpanError(span, err)
		return nil, handling.Process(err)
	}

	log.Info("submissions successfully founded")
	return foundFormSubmissions, nil
}
