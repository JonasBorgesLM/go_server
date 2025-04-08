// Package dto contains data transfer objects used in the API.
package dto

import "github.com/JonasBorgesLM/go_server/pkg/validations"

// UserResponse represents the data returned when querying a user.
type UserResponse struct {
	ID       string `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CreateUserResponse represents the data returned after creating a user.
type CreateUserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

// CreateUserRequest represents the data required to create a new user.
type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,strong_password"`
	Username string `json:"username" validate:"required,username"`
}

// Validate checks if the CreateUserRequest fields meet the required criteria.
func (cur *CreateUserRequest) Validate() error {
	return validations.Validate.Struct(cur)
}

// UpdateUserRequest represents the data required to update an existing user.
type UpdateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,username"`
}

// Validate checks if the UpdateUserRequest fields meet the required criteria.
func (uur *UpdateUserRequest) Validate() error {
	return validations.Validate.Struct(uur)
}
