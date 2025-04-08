package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidateUsername checks if the provided username contains only alphabetic characters and spaces.
func ValidateUsername(fl validator.FieldLevel) bool {
	username := fl.Field().String()
	return regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString(username)
}
