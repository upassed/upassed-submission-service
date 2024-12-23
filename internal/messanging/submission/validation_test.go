package submission_test

import (
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-answer-service/internal/util"
	"testing"
)

func TestSubmissionCreateRequestValidation_InvalidFormID(t *testing.T) {
	request := util.RandomEventSubmissionCreateRequest()
	request.FormID = "invalid_uuid"

	err := request.Validate()
	require.Error(t, err)
}

func TestSubmissionCreateRequestValidation_InvalidQuestionID(t *testing.T) {
	request := util.RandomEventSubmissionCreateRequest()
	request.QuestionID = "invalid_uuid"

	err := request.Validate()
	require.Error(t, err)
}

func TestSubmissionCreateRequestValidation_InvalidAnswerID(t *testing.T) {
	request := util.RandomEventSubmissionCreateRequest()
	request.AnswerIDs[0] = "invalid_uuid"

	err := request.Validate()
	require.Error(t, err)
}

func TestSubmissionCreateRequestValidation_EmptyAnswerIDs(t *testing.T) {
	request := util.RandomEventSubmissionCreateRequest()
	request.AnswerIDs = nil

	err := request.Validate()
	require.Error(t, err)
}
