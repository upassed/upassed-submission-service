package submission

import (
	"context"
	"errors"
	"github.com/upassed/upassed-submission-service/internal/handling"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
	"github.com/upassed/upassed-submission-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

var (
	ErrStudentFormSubmissionsNotFound  = errors.New("student's form submissions were not found in the database")
	errSearchingStudentFormSubmissions = errors.New("error while searching student's form submissions")
)

func (repository *repositoryImpl) FindStudentFormSubmissions(ctx context.Context, params *domain.StudentFormSubmissionsSearchParams) ([]*domain.Submission, error) {
	username := ctx.Value(auth.UsernameKey).(string)

	_, span := otel.Tracer(repository.cfg.Tracing.SubmissionTracerName).Start(ctx, "submissionRepository#FindStudentFormSubmissions")
	span.SetAttributes(
		attribute.String("username", username),
		attribute.String("formID", params.FormID.String()),
	)
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.FindStudentFormSubmissions),
		logging.WithCtx(ctx),
		logging.WithAny("formID", params.FormID.String()),
	)

	log.Info("started searching student's submissions in the database")
	foundSubmissions := make([]*domain.Submission, 0)
	findResult := repository.db.WithContext(ctx).Model(&domain.Submission{}).
		Where("form_id = ?", params.FormID).
		Where("student_username = ?", params.StudentUsername).
		Find(&foundSubmissions)

	if err := findResult.Error; err != nil {
		log.Error("error while searching student's submissions", logging.Error(err))
		tracing.SetSpanError(span, err)
		return nil, handling.New(errSearchingStudentFormSubmissions.Error(), codes.Internal)
	}

	if len(foundSubmissions) == 0 {
		log.Error("student's submissions were not found in the database")
		tracing.SetSpanError(span, ErrStudentFormSubmissionsNotFound)
		return nil, handling.New(ErrStudentFormSubmissionsNotFound.Error(), codes.NotFound)
	}

	log.Info("student's form submissions were successfully found in the database")
	return foundSubmissions, nil
}
