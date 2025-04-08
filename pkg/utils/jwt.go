// Package utils provides utility functions for generating JSON Web Tokens (JWT) using HMAC-SHA256 for authentication purposes.
package utils

import (
	"errors"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// GenerateJWT creates a JSON Web Token (JWT) for the given user ID, signed with HMAC-SHA256 and valid for 24 hours.
func GenerateJWT(userID string) (string, error) {
	jwtKey := []byte(os.Getenv("JWT_SECRET"))
	if jwtKey == nil {
		return "", errors.New("JWT_SECRET not configured")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
