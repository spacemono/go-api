package validator

import "strings"

type ValidationError struct {
	details map[string][]string
}

func NewValidationError(details map[string][]string) error {
	return &ValidationError{details: details}
}

func (ve *ValidationError) Error() string {
	var formattedErrors []string

	for field, issues := range ve.details {
		formattedErrors = append(formattedErrors, field+" is "+strings.Join(issues, ", "))
	}

	return "validation error: " + strings.Join(formattedErrors, "; ")
}
