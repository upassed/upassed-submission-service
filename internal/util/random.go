package util

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	event "github.com/upassed/upassed-submission-service/internal/messanging/model"
	domain "github.com/upassed/upassed-submission-service/internal/repository/model"
	business "github.com/upassed/upassed-submission-service/internal/service/model"
)

func RandomEventSubmissionCreateRequest() *event.SubmissionCreateRequest {
	numberOfChosenAnswers := gofakeit.IntRange(1, 4)
	answerIDs := make([]string, 0, numberOfChosenAnswers)
	for i := 0; i < numberOfChosenAnswers; i++ {
		answerIDs = append(answerIDs, uuid.NewString())
	}

	return &event.SubmissionCreateRequest{
		FormID:     uuid.NewString(),
		QuestionID: uuid.NewString(),
		AnswerIDs:  answerIDs,
	}
}

func RandomBusinessSubmission() *business.Submission {
	numberOfChosenAnswers := gofakeit.IntRange(1, 4)
	answerIDs := make([]uuid.UUID, 0, numberOfChosenAnswers)
	for i := 0; i < numberOfChosenAnswers; i++ {
		answerIDs = append(answerIDs, uuid.New())
	}

	return &business.Submission{
		StudentUsername: gofakeit.Username(),
		FormID:          uuid.New(),
		QuestionID:      uuid.New(),
		AnswerIDs:       answerIDs,
	}
}

func RandomDomainSubmissions() []*domain.Submission {
	numberOfSubmissions := gofakeit.IntRange(2, 10)
	submissions := make([]*domain.Submission, 0, numberOfSubmissions)

	formID := uuid.New()
	questionID := uuid.New()
	studentUsername := gofakeit.Username()

	for i := 0; i < numberOfSubmissions; i++ {
		submissions = append(submissions, &domain.Submission{
			ID:              uuid.New(),
			StudentUsername: studentUsername,
			FormID:          formID,
			QuestionID:      questionID,
			AnswerID:        uuid.New(),
		})
	}

	return submissions
}

func RandomBusinessFormSubmissions() *business.FormSubmissions {
	questionSubmissionsNumber := gofakeit.IntRange(10, 20)
	questionSubmissions := make([]*business.QuestionSubmission, 0, questionSubmissionsNumber)

	for i := 0; i < questionSubmissionsNumber; i++ {
		answerIDsNumber := gofakeit.IntRange(1, 5)
		answerIDs := make([]uuid.UUID, 0, answerIDsNumber)

		for j := 0; j < answerIDsNumber; j++ {
			answerIDs = append(answerIDs, uuid.New())
		}

		questionSubmissions = append(questionSubmissions, &business.QuestionSubmission{
			QuestionID: uuid.New(),
			AnswerIDs:  answerIDs,
		})
	}

	return &business.FormSubmissions{
		StudentUsername:     gofakeit.Username(),
		FormID:              uuid.New(),
		QuestionSubmissions: questionSubmissions,
	}
}
