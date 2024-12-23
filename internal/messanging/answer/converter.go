package answer

import (
	"encoding/json"
	"github.com/google/uuid"
	event "github.com/upassed/upassed-answer-service/internal/messanging/model"
	business "github.com/upassed/upassed-answer-service/internal/service/model"
)

func ConvertToAnswerCreateRequest(messageBody []byte) (*event.AnswerCreateRequest, error) {
	var request event.AnswerCreateRequest
	if err := json.Unmarshal(messageBody, &request); err != nil {
		return nil, err
	}

	return &request, nil
}

func ConvertToBusinessAnswer(request *event.AnswerCreateRequest, studentUsername string) *business.Answer {
	answerIDs := make([]uuid.UUID, 0, len(request.AnswerIDs))
	for _, answerID := range request.AnswerIDs {
		answerIDs = append(answerIDs, uuid.MustParse(answerID))
	}

	return &business.Answer{
		ID:              uuid.New(),
		StudentUsername: studentUsername,
		FormID:          uuid.MustParse(request.FormID),
		QuestionID:      uuid.MustParse(request.QuestionID),
		AnswerIDs:       answerIDs,
	}
}
