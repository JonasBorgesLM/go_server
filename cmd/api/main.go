// Package main is the entry point for the API server application.
package main

import (
	"log"

	"github.com/JonasBorgesLM/go_server/internal/config"
	"github.com/JonasBorgesLM/go_server/internal/server"
)

func main() {
	if err := config.LoadEnv(".env"); err != nil {
		log.Fatal("Failed to load environment variables: ", err)
	}

	if err := config.ValidateEnv(); err != nil {
		log.Fatal("Failed to validate environment variables: ", err)
	}

	srv := server.NewServer()
	if err := srv.Start(); err != nil {
		log.Fatal("Failed to start the server: ", err)
	}
}
