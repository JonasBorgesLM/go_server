// Package validations provides custom validation functions for various data types, such as email addresses.
package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidateCustomEmail checks if the provided email address matches a standard email format using a regular expression.
func ValidateCustomEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`).MatchString(email)
}
