package event

import (
	"github.com/go-playground/validator/v10"
)

type AnswerCreateRequest struct {
	FormID     string   `json:"form_id,omitempty" validate:"required,uuid"`
	QuestionID string   `json:"question_id,omitempty" validate:"required,uuid"`
	AnswerIDs  []string `json:"answer_id,omitempty" validate:"required,uuid"`
}

func (request *AnswerCreateRequest) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("uuid", ValidateUUID())

	if err := validate.Struct(*request); err != nil {
		return err
	}

	return nil
}
