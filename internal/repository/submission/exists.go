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
	errCountingSubmissions = errors.New("")
)

func (repository *repositoryImpl) Exists(ctx context.Context, params *domain.SubmissionExistCheckParams) (bool, error) {
	studentUsername := ctx.Value(auth.UsernameKey).(string)

	_, span := otel.Tracer(repository.cfg.Tracing.SubmissionTracerName).Start(ctx, "submissionRepository#Exists")
	span.SetAttributes(
		attribute.String("studentUsername", studentUsername),
		attribute.String("formID", params.FormID.String()),
		attribute.String("questionID", params.QuestionID.String()),
	)
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.Exists),
		logging.WithCtx(ctx),
		logging.WithAny("formID", params.FormID.String()),
		logging.WithAny("questionID", params.QuestionID.String()),
	)

	log.Info("started searching existing submissions in the database")
	var submissionsCount int64
	countResult := repository.db.WithContext(ctx).Model(&domain.Submission{}).
		Where("form_id = ?", params.FormID).
		Where("question_id = ?", params.QuestionID).
		Where("student_username = ?", studentUsername).
		Count(&submissionsCount)

	if err := countResult.Error; err != nil {
		log.Error("error while counting existing submissions", logging.Error(err))
		tracing.SetSpanError(span, err)
		return false, handling.New(errCountingSubmissions.Error(), codes.Internal)
	}

	if submissionsCount > 0 {
		log.Info("found existing submissions in the database")
		return true, nil
	}

	log.Info("existing submissions not found in the database")
	return false, nil
}
