package submission

import (
	"encoding/json"
	"github.com/google/uuid"
	event "github.com/upassed/upassed-answer-service/internal/messanging/model"
	business "github.com/upassed/upassed-answer-service/internal/service/model"
)

func ConvertToSubmissionCreateRequest(messageBody []byte) (*event.SubmissionCreateRequest, error) {
	var request event.SubmissionCreateRequest
	if err := json.Unmarshal(messageBody, &request); err != nil {
		return nil, err
	}

	return &request, nil
}

func ConvertToBusinessSubmission(request *event.SubmissionCreateRequest, studentUsername string) *business.Submission {
	answerIDs := make([]uuid.UUID, 0, len(request.AnswerIDs))
	for _, answerID := range request.AnswerIDs {
		answerIDs = append(answerIDs, uuid.MustParse(answerID))
	}

	return &business.Submission{
		StudentUsername: studentUsername,
		FormID:          uuid.MustParse(request.FormID),
		QuestionID:      uuid.MustParse(request.QuestionID),
		AnswerIDs:       answerIDs,
	}
}
