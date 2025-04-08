// Package formatter provides utility functions to convert between domain models and DTOs.
package formatter

import (
	"github.com/JonasBorgesLM/go_server/internal/domain"
	"github.com/JonasBorgesLM/go_server/internal/dto"
	"golang.org/x/crypto/bcrypt"
)

// UserToCreateUserResponse converts a domain User to a CreateUserResponse DTO.
func UserToCreateUserResponse(user *domain.User) *dto.CreateUserResponse {
	return &dto.CreateUserResponse{
		Username: user.Username,
		Email:    user.Email,
	}
}

// CreateUserRequestToDomain converts a CreateUserRequest DTO to a domain User.
func CreateUserRequestToDomain(request dto.CreateUserRequest) (*domain.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &domain.User{
		Email:    request.Email,
		Username: request.Username,
		Password: string(hashedPassword),
	}, nil
}
