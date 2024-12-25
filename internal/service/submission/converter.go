package submission

import (
	"github.com/google/uuid"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
)

func ConvertToDomainSubmissions(businessSubmission *business.Submission) []*domain.Submission {
	domainSubmissions := make([]*domain.Submission, 0, len(businessSubmission.AnswerIDs))

	for _, answerID := range businessSubmission.AnswerIDs {
		domainSubmissions = append(domainSubmissions, &domain.Submission{
			ID:              uuid.New(),
			StudentUsername: businessSubmission.StudentUsername,
			FormID:          businessSubmission.FormID,
			QuestionID:      businessSubmission.QuestionID,
			AnswerID:        answerID,
		})
	}

	return domainSubmissions
}

func ConvertToSubmissionCreateResponse(domainSubmissions []*domain.Submission) *business.SubmissionCreateResponse {
	createdSubmissionIDs := make([]uuid.UUID, 0, len(domainSubmissions))

	for _, domainSubmission := range domainSubmissions {
		createdSubmissionIDs = append(createdSubmissionIDs, domainSubmission.ID)
	}

	return &business.SubmissionCreateResponse{
		CreatedSubmissionIDs: createdSubmissionIDs,
	}
}

func ConvertToSubmissionExistCheckParams(submission *business.Submission) *domain.SubmissionExistCheckParams {
	return &domain.SubmissionExistCheckParams{
		StudentUsername: submission.StudentUsername,
		FormID:          submission.FormID,
		QuestionID:      submission.QuestionID,
	}
}

func ConvertToSubmissionDeleteParams(submission *business.Submission) *domain.SubmissionDeleteParams {
	return &domain.SubmissionDeleteParams{
		StudentUsername: submission.StudentUsername,
		FormID:          submission.FormID,
		QuestionID:      submission.QuestionID,
	}
}
