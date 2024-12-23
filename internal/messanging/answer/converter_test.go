package answer_test

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-answer-service/internal/messanging/answer"
	"github.com/upassed/upassed-answer-service/internal/util"
	"testing"
)

func TestConvertToAnswerCreateRequest_InvalidBytes(t *testing.T) {
	invalidBytes := make([]byte, 10)
	_, err := answer.ConvertToAnswerCreateRequest(invalidBytes)
	require.Error(t, err)
}

func TestConvertToAnswerCreateRequest_ValidBytes(t *testing.T) {
	initialRequest := util.RandomEventAnswerCreateRequest()
	initialRequestBytes, err := json.Marshal(initialRequest)
	require.NoError(t, err)

	convertedRequest, err := answer.ConvertToAnswerCreateRequest(initialRequestBytes)
	require.NoError(t, err)

	assert.Equal(t, initialRequest.FormID, convertedRequest.FormID)
	assert.Equal(t, initialRequest.QuestionID, convertedRequest.QuestionID)
	assert.Equal(t, len(initialRequest.AnswerIDs), len(initialRequest.AnswerIDs))

	for idx, answerID := range initialRequest.AnswerIDs {
		assert.Equal(t, answerID, convertedRequest.AnswerIDs[idx])
	}
}

func TestConvertToBusinessAnswer(t *testing.T) {
	studentUsername := gofakeit.Username()
	answerCreateRequest := util.RandomEventAnswerCreateRequest()

	businessAnswer := answer.ConvertToBusinessAnswer(answerCreateRequest, studentUsername)
	assert.NotNil(t, businessAnswer.ID)
	assert.Equal(t, answerCreateRequest.FormID, businessAnswer.FormID.String())
	assert.Equal(t, answerCreateRequest.QuestionID, businessAnswer.QuestionID.String())
	assert.Equal(t, studentUsername, businessAnswer.StudentUsername)
	assert.Equal(t, len(answerCreateRequest.AnswerIDs), len(businessAnswer.AnswerIDs))

	for idx, answerID := range answerCreateRequest.AnswerIDs {
		assert.Equal(t, answerID, businessAnswer.AnswerIDs[idx].String())
	}
}
