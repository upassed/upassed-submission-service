package submission_test

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-submission-service/internal/messanging/submission"
	"github.com/upassed/upassed-submission-service/internal/util"
	"testing"
)

func TestConvertToSubmissionCreateRequest_InvalidBytes(t *testing.T) {
	invalidBytes := make([]byte, 10)
	_, err := submission.ConvertToSubmissionCreateRequest(invalidBytes)
	require.Error(t, err)
}

func TestConvertToSubmissionCreateRequest_ValidBytes(t *testing.T) {
	initialRequest := util.RandomEventSubmissionCreateRequest()
	initialRequestBytes, err := json.Marshal(initialRequest)
	require.NoError(t, err)

	convertedRequest, err := submission.ConvertToSubmissionCreateRequest(initialRequestBytes)
	require.NoError(t, err)

	assert.Equal(t, initialRequest.FormID, convertedRequest.FormID)
	assert.Equal(t, initialRequest.QuestionID, convertedRequest.QuestionID)
	assert.Equal(t, len(initialRequest.AnswerIDs), len(initialRequest.AnswerIDs))

	for idx, answerID := range initialRequest.AnswerIDs {
		assert.Equal(t, answerID, convertedRequest.AnswerIDs[idx])
	}
}

func TestConvertToBusinessSubmission(t *testing.T) {
	studentUsername := gofakeit.Username()
	answerCreateRequest := util.RandomEventSubmissionCreateRequest()

	businessAnswer := submission.ConvertToBusinessSubmission(answerCreateRequest, studentUsername)
	assert.Equal(t, answerCreateRequest.FormID, businessAnswer.FormID.String())
	assert.Equal(t, answerCreateRequest.QuestionID, businessAnswer.QuestionID.String())
	assert.Equal(t, studentUsername, businessAnswer.StudentUsername)
	assert.Equal(t, len(answerCreateRequest.AnswerIDs), len(businessAnswer.AnswerIDs))

	for idx, answerID := range answerCreateRequest.AnswerIDs {
		assert.Equal(t, answerID, businessAnswer.AnswerIDs[idx].String())
	}
}
