// Package config_test contains tests for validating environment variables in the configuration.
package config_test

import (
	"log"
	"os"

	"github.com/JonasBorgesLM/go_server/internal/config"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("ValidateEnv", func() {
	ginkgo.AfterEach(func() {
		clearEnv()
	})

	tests := []struct {
		name          string
		envVars       map[string]string
		expectedError string
	}{
		{
			name: "valid environment variables",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USERNAME": "user",
				"DB_PASSWORD": "password",
				"DB_DATABASE": "dbname",
				"JWT_SECRET":  "secret",
				"PORT":        "8080",
			},
			expectedError: "",
		},
		{
			name: "invalid DB_PORT",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "not_a_number",
				"DB_USERNAME": "user",
				"DB_PASSWORD": "password",
				"DB_DATABASE": "dbname",
				"JWT_SECRET":  "secret",
				"PORT":        "8080",
			},
			expectedError: "DB_PORT must be a valid number",
		},
		{
			name: "empty DB_USERNAME",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USERNAME": "",
				"DB_PASSWORD": "password",
				"DB_DATABASE": "dbname",
				"JWT_SECRET":  "secret",
				"PORT":        "8080",
			},
			expectedError: "DB_USERNAME cannot be empty",
		},
		{
			name: "empty DB_PASSWORD",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USERNAME": "user",
				"DB_PASSWORD": "",
				"DB_DATABASE": "dbname",
				"JWT_SECRET":  "secret",
				"PORT":        "8080",
			},
			expectedError: "DB_PASSWORD cannot be empty",
		},
		{
			name: "invalid PORT",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USERNAME": "user",
				"DB_PASSWORD": "password",
				"DB_DATABASE": "dbname",
				"JWT_SECRET":  "secret",
				"PORT":        "not_a_number",
			},
			expectedError: "PORT must be a valid number",
		},
		{
			name: "empty DB_DATABASE",
			envVars: map[string]string{
				"DB_HOST":     "localhost",
				"DB_PORT":     "5432",
				"DB_USERNAME": "user",
				"DB_PASSWORD": "password",
				"DB_DATABASE": "",
				"JWT_SECRET":  "secret",
				"PORT":        "8080",
			},
			expectedError: "DB_DATABASE cannot be empty",
		},
	}

	for _, tt := range tests {
		ginkgo.Context(tt.name, func() {
			ginkgo.BeforeEach(func() {
				setEnvVars(tt.envVars)
			})

			ginkgo.It("should validate environment variables correctly", func() {
				err := config.ValidateEnv()
				if tt.expectedError != "" {
					gomega.Expect(err).To(gomega.HaveOccurred())
					gomega.Expect(err.Error()).To(gomega.Equal(tt.expectedError))
				} else {
					gomega.Expect(err).To(gomega.BeNil())
				}
			})
		})
	}
})

// Function to clear environment variables
func clearEnv() {
	envVars := []string{
		"TEST_KEY",
		"DB_HOST",
		"DB_PORT",
		"DB_USERNAME",
		"DB_PASSWORD",
		"DB_DATABASE",
		"JWT_SECRET",
		"PORT",
	}
	for _, v := range envVars {
		if err := os.Unsetenv(v); err != nil {
			log.Printf("Error unsetting environment variable %s: %v", v, err)
		}
	}
}

// Function to set environment variables
func setEnvVars(vars map[string]string) {
	for key, value := range vars {
		if err := os.Setenv(key, value); err != nil {
			log.Printf("Error setting environment variable %s: %v", key, err)
		}
	}
}
