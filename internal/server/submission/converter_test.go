package submission_test

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/upassed/upassed-submission-service/internal/server/submission"
	"github.com/upassed/upassed-submission-service/internal/util"
	"testing"
)

func TestConvertToFindStudentFormSubmissionsResponse(t *testing.T) {
	formSubmissions := util.RandomBusinessFormSubmissions()
	convertedResponse := submission.ConvertToFindStudentFormSubmissionsResponse(formSubmissions)

	assert.Equal(t, formSubmissions.StudentUsername, convertedResponse.GetStudentUsername())
	assert.Equal(t, formSubmissions.FormID.String(), convertedResponse.GetFormId())
	assert.Equal(t, len(formSubmissions.QuestionSubmissions), len(convertedResponse.GetQuestionSubmissions()))

	for questionIdx, questionSubmission := range formSubmissions.QuestionSubmissions {
		convertedQuestionSubmission := convertedResponse.GetQuestionSubmissions()[questionIdx]

		assert.Equal(t, questionSubmission.QuestionID.String(), convertedQuestionSubmission.GetQuestionId())
		assert.Equal(t, len(questionSubmission.AnswerIDs), len(convertedQuestionSubmission.GetAnswerIds()))

		for answerIdx, answerID := range questionSubmission.AnswerIDs {
			assert.Equal(t, answerID.String(), convertedQuestionSubmission.GetAnswerIds()[answerIdx])
		}
	}
}

func TestConvertToStudentFormSubmissionsSearchParams(t *testing.T) {
	studentUsername := gofakeit.Username()
	formID := uuid.New()

	params := submission.ConvertToStudentFormSubmissionsSearchParams(studentUsername, formID)

	assert.Equal(t, studentUsername, params.StudentUsername)
	assert.Equal(t, formID, params.FormID)
}
