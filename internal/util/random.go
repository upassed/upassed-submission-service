package util

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	event "github.com/upassed/upassed-answer-service/internal/messanging/model"
)

func RandomEventAnswerCreateRequest() *event.AnswerCreateRequest {
	numberOfChosenAnswers := gofakeit.IntRange(1, 4)
	answerIDs := make([]string, 0, numberOfChosenAnswers)
	for i := 0; i < numberOfChosenAnswers; i++ {
		answerIDs = append(answerIDs, uuid.NewString())
	}

	return &event.AnswerCreateRequest{
		FormID:     uuid.NewString(),
		QuestionID: uuid.NewString(),
		AnswerIDs:  answerIDs,
	}
}
