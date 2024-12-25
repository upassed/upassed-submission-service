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
	errSavingSubmissions = errors.New("error while saving submissions into the database")
)

func (repository *repositoryImpl) Save(ctx context.Context, submissions []*domain.Submission) error {
	studentUsername := ctx.Value(auth.UsernameKey).(string)

	_, span := otel.Tracer(repository.cfg.Tracing.SubmissionTracerName).Start(ctx, "submissionRepository#Save")
	span.SetAttributes(
		attribute.String("studentUsername", studentUsername),
		attribute.String("formID", submissions[0].FormID.String()),
		attribute.String("questionID", submissions[0].QuestionID.String()),
		attribute.Int("answersCount", len(submissions)),
	)
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.Save),
		logging.WithCtx(ctx),
		logging.WithAny("formID", submissions[0].FormID.String()),
		logging.WithAny("questionID", submissions[0].QuestionID.String()),
		logging.WithAny("answersCount", len(submissions)),
	)

	log.Info("started saving submissions into the database")
	saveResult := repository.db.WithContext(ctx).CreateInBatches(submissions, 100)
	if err := saveResult.Error; err != nil {
		log.Error("error while saving submissions into the database", logging.Error(err))
		tracing.SetSpanError(span, err)
		return handling.New(errSavingSubmissions.Error(), codes.Internal)
	}

	log.Info("submissions were successfully inserted into the database")
	return nil
}
