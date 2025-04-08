// Package config_test contains tests for loading environment variables from a .env file.
package config_test

import (
	"log"
	"os"

	"github.com/JonasBorgesLM/go_server/internal/config"
	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
)

var _ = ginkgo.Describe("LoadEnv", func() {
	ginkgo.It("should load environment variables from a .env file", func() {
		file, err := os.CreateTemp("", ".env")
		gomega.Expect(err).To(gomega.BeNil())
		defer func() {
			if err := os.Remove(file.Name()); err != nil {
				log.Printf("Error removing temp file: %v", err)
			}
		}()

		_, err = file.WriteString("DB_HOST=localhost\nDB_PORT=5432\nDB_USERNAME=user\nDB_PASSWORD=password\nDB_DATABASE=dbname\nJWT_SECRET=secret\nPORT=8080\n")
		gomega.Expect(err).To(gomega.BeNil())

		err = file.Close()
		gomega.Expect(err).To(gomega.BeNil())

		err = config.LoadEnv(file.Name())
		gomega.Expect(err).To(gomega.BeNil())

		gomega.Expect(os.Getenv("DB_HOST")).To(gomega.Equal("localhost"))
		gomega.Expect(os.Getenv("DB_PORT")).To(gomega.Equal("5432"))
	})

	ginkgo.It("should return an error if the .env file cannot be loaded", func() {
		err := config.LoadEnv("non_existent_file.env")
		gomega.Expect(err).To(gomega.HaveOccurred())
		gomega.Expect(err.Error()).To(gomega.ContainSubstring("Error loading .env file:"))
	})
})
