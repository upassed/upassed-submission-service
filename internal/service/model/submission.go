package business

import "github.com/google/uuid"

type Submission struct {
	StudentUsername string
	FormID          uuid.UUID
	QuestionID      uuid.UUID
	AnswerIDs       uuid.UUIDs
}

type SubmissionCreateResponse struct {
	CreatedSubmissionIDs uuid.UUIDs
}
