package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func FormatValidationError(err error) []string {
	var errors []string
	if _, ok := err.(*validator.InvalidValidationError); ok {
		errors = append(errors, "Validation failed")
		return errors
	}

	for _, err := range err.(validator.ValidationErrors) {
		switch err.Tag() {
		case "required":
			errors = append(errors, fmt.Sprintf("%s is required", err.Field()))
		case "email":
			errors = append(errors, fmt.Sprintf("%s must be a valid email address", err.Field()))
		case "min":
			errors = append(errors, fmt.Sprintf("%s must be at least %s characters long", err.Field(), err.Param()))
		default:
			errors = append(errors, fmt.Sprintf("Validation failed for %s", err.Field()))
		}
	}
	return errors
}
