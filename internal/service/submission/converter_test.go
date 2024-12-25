package submission_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/upassed/upassed-submission-service/internal/service/submission"
	"github.com/upassed/upassed-submission-service/internal/util"
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
