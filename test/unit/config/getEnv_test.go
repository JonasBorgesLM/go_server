// Package config_test contains tests for validating environment variables in the configuration.
package config_test

import (
	"log"
	"os"

	"github.com/JonasBorgesLM/go_server/internal/config"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("GetEnv", func() {
	ginkgo.It("should return the value of an existing environment variable", func() {
		if err := os.Setenv("TEST_KEY", "test_value"); err != nil {
			log.Printf("Error setting environment variable: %v", err)
		}

		value, err := config.GetEnv("TEST_KEY")
		gomega.Expect(err).To(gomega.BeNil())
		gomega.Expect(value).To(gomega.Equal("test_value"))
	})

	ginkgo.It("should return an error for a non-existent environment variable", func() {
		_, err := config.GetEnv("NON_EXISTENT_KEY")
		gomega.Expect(err).To(gomega.HaveOccurred())
		gomega.Expect(err.Error()).To(gomega.Equal("environment variable not found: NON_EXISTENT_KEY"))
	})
})
