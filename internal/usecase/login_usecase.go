// Package usecase provides business logic implementations for login application.
package usecase

import (
	"errors"

	"github.com/JonasBorgesLM/go_server/pkg/utils"
)

// Execute handles the login process by validating the user's credentials and generating a JWT token.
func (uc *UserUseCase) Execute(email, password string) (string, error) {
	user, err := uc.repo.FindByEmail(email)
	if err != nil {
		return "", errors.New("user not found")
	}

	if !utils.ComparePasswords(user.Password, password) {
		return "", errors.New("incorrect password")
	}

	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		return "", errors.New("error generating token")
	}

	return token, nil
}
