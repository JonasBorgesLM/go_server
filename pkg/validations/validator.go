package validations

import (
	"log"

	"github.com/go-playground/validator/v10"
)

// Validate is a global instance of validator.Validate used to register and perform custom validations.
var Validate *validator.Validate

func init() {
	Validate = validator.New()

	registerValidation("custom_email", ValidateCustomEmail)
	registerValidation("strong_password", ValidateStrongPassword)
	registerValidation("username", ValidateUsername)
}

func registerValidation(tag string, fn validator.Func) {
	if err := Validate.RegisterValidation(tag, fn); err != nil {
		log.Fatalf("Error registering validation '%s': %v", tag, err)
	}
}
