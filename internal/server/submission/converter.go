package submission

import (
	"github.com/google/uuid"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
	"github.com/upassed/upassed-submission-service/pkg/client"
)

func ConvertToFindStudentFormSubmissionsResponse(foundSubmissions *business.FormSubmissions) *client.FindStudentFormSubmissionsResponse {
	questionSubmissions := make([]*client.QuestionSubmission, 0, len(foundSubmissions.QuestionSubmissions))
	for _, questionSubmission := range foundSubmissions.QuestionSubmissions {
		answerIDs := make([]string, 0, len(questionSubmission.AnswerIDs))
		for _, answerID := range questionSubmission.AnswerIDs {
			answerIDs = append(answerIDs, answerID.String())
		}

		questionSubmissions = append(questionSubmissions, &client.QuestionSubmission{
			QuestionId: questionSubmission.QuestionID.String(),
			AnswerIds:  answerIDs,
		})
	}

	return &client.FindStudentFormSubmissionsResponse{
		StudentUsername:     foundSubmissions.StudentUsername,
		FormId:              foundSubmissions.FormID.String(),
		QuestionSubmissions: questionSubmissions,
	}
}

func ConvertToStudentFormSubmissionsSearchParams(studentUsername string, formID uuid.UUID) *business.StudentFormSubmissionSearchParams {
	return &business.StudentFormSubmissionSearchParams{
		StudentUsername: studentUsername,
		FormID:          formID,
	}
}
