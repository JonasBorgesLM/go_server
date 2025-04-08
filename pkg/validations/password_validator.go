package validations

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

// ValidateStrongPassword checks if the provided password meets strong password criteria.
func ValidateStrongPassword(fl validator.FieldLevel) bool {
	password := fl.Field().String()
	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*()_+{}:"<>?]`).MatchString(password)
	return len(password) >= 8 && hasUpper && hasLower && hasNumber && hasSpecial
}
