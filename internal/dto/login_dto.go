// Package dto contains data transfer objects used in the login application.
package dto

import "github.com/JonasBorgesLM/go_server/pkg/validations"

// LoginRequest represents the data required for a user login.
type LoginRequest struct {
	Email    string `json:"email" validate:"required,custom_email"`
	Password string `json:"password" validate:"required,strong_password"`
}

// Validate checks if the LoginRequest fields meet the required criteria.
func (lr *LoginRequest) Validate() error {
	return validations.Validate.Struct(lr)
}
