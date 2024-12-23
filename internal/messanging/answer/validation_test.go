package answer_test

import (
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-answer-service/internal/util"
	"testing"
)

func TestAnswerCreateRequestValidation_InvalidFormID(t *testing.T) {
	request := util.RandomEventAnswerCreateRequest()
	request.FormID = "invalid_uuid"

	err := request.Validate()
	require.Error(t, err)
}

func TestAnswerCreateRequestValidation_InvalidQuestionID(t *testing.T) {
	request := util.RandomEventAnswerCreateRequest()
	request.QuestionID = "invalid_uuid"

	err := request.Validate()
	require.Error(t, err)
}

func TestAnswerCreateRequestValidation_InvalidAnswerID(t *testing.T) {
	request := util.RandomEventAnswerCreateRequest()
	request.AnswerIDs[0] = "invalid_uuid"

	err := request.Validate()
	require.Error(t, err)
}

func TestAnswerCreateRequestValidation_EmptyAnswerIDs(t *testing.T) {
	request := util.RandomEventAnswerCreateRequest()
	request.AnswerIDs = nil

	err := request.Validate()
	require.Error(t, err)
}
