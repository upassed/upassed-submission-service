package domain

import "github.com/google/uuid"

type Submission struct {
	ID              uuid.UUID
	StudentUsername string
	FormID          uuid.UUID
	QuestionID      uuid.UUID
	AnswerID        uuid.UUID
}
