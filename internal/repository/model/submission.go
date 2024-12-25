package domain

import "github.com/google/uuid"

type Submission struct {
	ID              uuid.UUID
	StudentUsername string
	FormID          uuid.UUID
	QuestionID      uuid.UUID
	AnswerID        uuid.UUID
}

type SubmissionExistCheckParams struct {
	StudentUsername string
	FormID          uuid.UUID
	QuestionID      uuid.UUID
}

type SubmissionDeleteParams struct {
	StudentUsername string
	FormID          uuid.UUID
	QuestionID      uuid.UUID
}
