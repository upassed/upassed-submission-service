package submission

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-submission-service/internal/handling"
	"github.com/upassed/upassed-submission-service/internal/middleware/common/auth"
	requestid "github.com/upassed/upassed-submission-service/internal/middleware/common/request_id"
	"github.com/upassed/upassed-submission-service/internal/tracing"
	"github.com/upassed/upassed-submission-service/pkg/client"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

func (server *submissionServerAPI) FindStudentFormSubmissions(ctx context.Context, request *client.FindStudentFormSubmissionsRequest) (*client.FindStudentFormSubmissionsResponse, error) {
	username := ctx.Value(auth.UsernameKey).(string)

	spanContext, span := otel.Tracer(server.cfg.Tracing.SubmissionTracerName).Start(ctx, "submission#FindStudentFormSubmissions")
	span.SetAttributes(
		attribute.String(auth.UsernameKey, username),
		attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)),
		attribute.String("formID", request.GetFormId()),
	)
	defer span.End()

	if err := request.Validate(); err != nil {
		tracing.SetSpanError(span, err)
		return nil, handling.Wrap(err, handling.WithCode(codes.InvalidArgument))
	}

	response, err := server.service.FindByFormID(spanContext, uuid.MustParse(request.GetFormId()))
	if err != nil {
		tracing.SetSpanError(span, err)
		return nil, err
	}

	return ConvertToFindStudentFormSubmissionsResponse(response), nil
}
