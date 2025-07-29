package validator

import (
	"github.com/go-playground/validator/v10"
)

type Validator struct {
	v10 *validator.Validate
}

func New() (*Validator, error) {
	v := &Validator{v10: validator.New()}

	return v, nil
}

func (v *Validator) ValidateStruct(s any) error {
	if err := v.v10.Struct(s); err != nil {
		errorMessages := make(map[string][]string)

		for _, e := range err.(validator.ValidationErrors) {
			field := e.StructField()
			errorMessages[field] = append(errorMessages[field], e.Tag())
		}

		return NewValidationError(errorMessages)
	}

	return nil
}
