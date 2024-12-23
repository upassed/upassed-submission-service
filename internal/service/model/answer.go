package business

import "github.com/google/uuid"

type Answer struct {
	ID              uuid.UUID
	StudentUsername string
	FormID          uuid.UUID
	QuestionID      uuid.UUID
	AnswerIDs       uuid.UUIDs
}

type AnswerCreateResponse struct {
	CreatedAnswerIDs uuid.UUIDs
}
