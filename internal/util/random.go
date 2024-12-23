package util

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	event "github.com/upassed/upassed-submission-service/internal/messanging/model"
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
