// Package usecase provides business logic implementations for user management.
package usecase

import (
	"errors"

	"github.com/JonasBorgesLM/go_server/internal/domain"
	"github.com/JonasBorgesLM/go_server/internal/dto"
)

// UserUseCase provides use case operations for user management.
type UserUseCase struct {
	repo domain.UserRepository
}

// NewUserUseCase creates a new instance of UserUseCase with the provided UserRepository.
func NewUserUseCase(repo domain.UserRepository) *UserUseCase {
	return &UserUseCase{repo}
}

// ListUsers retrieves a list of all users from the repository.
func (uc *UserUseCase) ListUsers() (*[]dto.UserResponse, error) {
	return uc.repo.List()
}

// CreateUser adds a new user to the repository after checking for duplicate emails.
func (uc *UserUseCase) CreateUser(user *domain.User) error {
	_, err := uc.repo.FindByEmail(user.Email)
	if err == nil {
		return errors.New("email already in use")
	}

	return uc.repo.Create(user)
}

// UpdateUser modifies an existing user in the repository based on the provided ID and update data.
func (uc *UserUseCase) UpdateUser(id string, user *dto.UpdateUserRequest) error {
	_, err := uc.repo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return uc.repo.Update(id, user)
}

// DeleteUser removes a user from the repository based on the provided ID.
func (uc *UserUseCase) DeleteUser(id string) error {
	_, err := uc.repo.FindByID(id)
	if err != nil {
		return errors.New("user not found")
	}

	return uc.repo.Delete(id)
}

// FindUserByUsername retrieves a user from the repository based on the provided username.
func (uc *UserUseCase) FindUserByUsername(username string) (*dto.UserResponse, error) {
	return uc.repo.FindByUsername(username)
}
