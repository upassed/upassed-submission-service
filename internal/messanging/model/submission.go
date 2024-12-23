package event

import (
	"github.com/go-playground/validator/v10"
)

type SubmissionCreateRequest struct {
	FormID     string   `json:"form_id,omitempty" validate:"required,uuid"`
	QuestionID string   `json:"question_id,omitempty" validate:"required,uuid"`
	AnswerIDs  []string `json:"answer_id,omitempty" validate:"required,uuid"`
}

func (request *SubmissionCreateRequest) Validate() error {
	validate := validator.New()
	_ = validate.RegisterValidation("uuid", ValidateUUID())

	if err := validate.Struct(*request); err != nil {
		return err
	}

	return nil
}
