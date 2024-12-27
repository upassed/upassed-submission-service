package submission_test

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
	"github.com/upassed/upassed-submission-service/internal/service/submission"
	"github.com/upassed/upassed-submission-service/internal/util"
	"slices"
	"testing"
)

func TestConvertToDomainSubmissions(t *testing.T) {
	businessSubmission := util.RandomBusinessSubmission()
	domainSubmissions := submission.ConvertToDomainSubmissions(businessSubmission)

	assert.Equal(t, len(businessSubmission.AnswerIDs), len(domainSubmissions))

	for idx, domainSubmission := range domainSubmissions {
		assert.NotNil(t, domainSubmission.ID)
		assert.Equal(t, businessSubmission.StudentUsername, domainSubmission.StudentUsername)
		assert.Equal(t, businessSubmission.FormID, domainSubmission.FormID)
		assert.Equal(t, businessSubmission.QuestionID, domainSubmission.QuestionID)
		assert.Equal(t, businessSubmission.AnswerIDs[idx], domainSubmission.AnswerID)
	}
}

func TestConvertToSubmissionCreateResponse(t *testing.T) {
	businessSubmission := util.RandomBusinessSubmission()
	domainSubmissions := submission.ConvertToDomainSubmissions(businessSubmission)

	submissionCreateResponse := submission.ConvertToSubmissionCreateResponse(domainSubmissions)
	assert.Equal(t, len(domainSubmissions), len(submissionCreateResponse.CreatedSubmissionIDs))

	for idx, domainSubmission := range domainSubmissions {
		assert.Equal(t, domainSubmission.ID, submissionCreateResponse.CreatedSubmissionIDs[idx])
	}
}

func TestConvertToSubmissionExistCheckParams(t *testing.T) {
	submissionToConvert := util.RandomBusinessSubmission()
	params := submission.ConvertToSubmissionExistCheckParams(submissionToConvert)

	assert.Equal(t, submissionToConvert.StudentUsername, params.StudentUsername)
	assert.Equal(t, submissionToConvert.FormID, params.FormID)
	assert.Equal(t, submissionToConvert.QuestionID, params.QuestionID)
}

func TestConvertToSubmissionDeleteParams(t *testing.T) {
	submissionToConvert := util.RandomBusinessSubmission()
	params := submission.ConvertToSubmissionDeleteParams(submissionToConvert)

	assert.Equal(t, submissionToConvert.StudentUsername, params.StudentUsername)
	assert.Equal(t, submissionToConvert.FormID, params.FormID)
	assert.Equal(t, submissionToConvert.QuestionID, params.QuestionID)
}

func TestConvertToFormSubmissions(t *testing.T) {
	questionAnswersMap := generateRandomQuestionAnswersMap()
	domainSubmissions := make([]*domain.Submission, 0, len(questionAnswersMap))

	studentUsername := gofakeit.Username()
	formID := uuid.New()
	for questionID, answerIDs := range questionAnswersMap {
		for _, answerID := range answerIDs {
			domainSubmissions = append(domainSubmissions, &domain.Submission{
				ID:              uuid.New(),
				StudentUsername: studentUsername,
				FormID:          formID,
				QuestionID:      questionID,
				AnswerID:        answerID,
			})
		}
	}

	formSubmissions := submission.ConvertToFormSubmissions(domainSubmissions)

	assert.Equal(t, studentUsername, formSubmissions.StudentUsername)
	assert.Equal(t, formID, formSubmissions.FormID)
	assert.Equal(t, len(questionAnswersMap), len(formSubmissions.QuestionSubmissions))

	totalAnswersCount := 0
	for _, questionSubmission := range formSubmissions.QuestionSubmissions {
		totalAnswersCount += len(questionSubmission.AnswerIDs)
		assert.True(t, slices.Equal(questionAnswersMap[questionSubmission.QuestionID], questionSubmission.AnswerIDs))
	}

	assert.Equal(t, len(domainSubmissions), totalAnswersCount)
}

func generateRandomQuestionAnswersMap() map[uuid.UUID]uuid.UUIDs {
	questionsNumber := gofakeit.IntRange(10, 40)
	resultMap := make(map[uuid.UUID]uuid.UUIDs, questionsNumber)
	for i := 0; i < questionsNumber; i++ {
		questionID := uuid.New()
		answersNumber := gofakeit.IntRange(2, 6)
		answerIDs := make(uuid.UUIDs, 0, answersNumber)
		for j := 0; j < answersNumber; j++ {
			answerIDs = append(answerIDs, uuid.New())
		}

		resultMap[questionID] = answerIDs
	}

	return resultMap
}
