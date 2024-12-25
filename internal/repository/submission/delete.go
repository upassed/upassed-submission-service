package submission

import (
	"context"
	"errors"
	"github.com/upassed/upassed-submission-service/internal/logging"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

var (
	errDeletingStudentExistingSubmissions = errors.New("unable to delete student's existing submissions")
)

func (repository *repositoryImpl) Delete(ctx context.Context, params *domain.SubmissionDeleteParams) error {
	studentUsername := ctx.Value(auth.UsernameKey).(string)

	_, span := otel.Tracer(repository.cfg.Tracing.SubmissionTracerName).Start(ctx, "submissionRepository#Delete")
	span.SetAttributes(
		attribute.String("studentUsername", studentUsername),
		attribute.String("formID", params.FormID.String()),
		attribute.String("questionID", params.QuestionID.String()),
	)
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.Delete),
		logging.WithCtx(ctx),
		logging.WithAny("formID", params.FormID.String()),
		logging.WithAny("questionID", params.QuestionID.String()),
	)

	log.Info("started deleting existing submissions in the database")
	deleteResult := repository.db.
		Where("form_id = ?", params.FormID).
		Where("question_id = ?", params.QuestionID).
		Where("student_username = ?", studentUsername).
		Delete(&domain.Submission{})

	if err := deleteResult.Error; err != nil {
		log.Error("unable to delete student's existing submissions to form question", logging.Error(err))
		return errDeletingStudentExistingSubmissions
	}

	log.Info("existing submissions were successfully deleted from the database")
	return nil
}
