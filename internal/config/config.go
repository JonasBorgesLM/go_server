// Package config handles environment variables and app settings.
package config

import (
	"errors"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// LoadEnv loads environment variables from a .env file.
func LoadEnv(filePath string) error {
	if err := godotenv.Load(filePath); err != nil {
		return errors.New("Error loading .env file: " + err.Error())
	}
	return nil
}

// GetEnv retrieves the value of an environment variable by its key.
func GetEnv(key string) (string, error) {
	value := os.Getenv(key)
	if value == "" {
		return "", errors.New("environment variable not found: " + key)
	}
	return value, nil
}

// ValidateEnv checks if all required environment variables are set and valid.
func ValidateEnv() error {
	requiredKeys := []string{
		"DB_HOST",
		"DB_PORT",
		"DB_USERNAME",
		"DB_PASSWORD",
		"DB_DATABASE",
		"JWT_SECRET",
		"PORT",
	}

	for _, key := range requiredKeys {
		value := os.Getenv(key)
		if value == "" {
			return errors.New(key + " cannot be empty")
		}
	}

	if _, err := strconv.Atoi(os.Getenv("DB_PORT")); err != nil {
		return errors.New("DB_PORT must be a valid number")
	}
	if _, err := strconv.Atoi(os.Getenv("PORT")); err != nil {
		return errors.New("PORT must be a valid number")
	}

	return nil
}
