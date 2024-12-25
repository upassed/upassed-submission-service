package business

import "github.com/google/uuid"

type Submission struct {
	StudentUsername string
	FormID          uuid.UUID
	QuestionID      uuid.UUID
	AnswerIDs       uuid.UUIDs
}

type FormSubmissions struct {
	StudentUsername     string
	FormID              uuid.UUID
	QuestionSubmissions []*QuestionSubmission
}

type QuestionSubmission struct {
	QuestionID uuid.UUID
	AnswerIDs  uuid.UUIDs
}

type SubmissionCreateResponse struct {
	CreatedSubmissionIDs uuid.UUIDs
}
