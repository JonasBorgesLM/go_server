// Package server provides the core functionality for running and managing the application's HTTP server.
package server

import (
	"log"
	"net/http"
	"os"

	"github.com/JonasBorgesLM/go_server/internal/handler"
	"github.com/JonasBorgesLM/go_server/internal/middleware"
	"github.com/JonasBorgesLM/go_server/internal/repository"
	"github.com/JonasBorgesLM/go_server/internal/usecase"
	"github.com/JonasBorgesLM/go_server/pkg/database"
)

// Server represents an HTTP server with a router and a configured port.
type Server struct {
	router *http.ServeMux
	port   string
}

// NewServer initializes and returns a new Server instance with configured routes and database connection.
func NewServer() *Server {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	userRepo := repository.NewUserRepository(db)
	userUseCase := usecase.NewUserUseCase(userRepo)
	userHandler := handler.NewUserHandler(userUseCase)

	router := http.NewServeMux()
	router.HandleFunc("/login", userHandler.Login)
	router.HandleFunc("/users", middleware.AuthMiddleware(userHandler.ListUsers))
	router.HandleFunc("/users/create", middleware.AuthMiddleware(userHandler.CreateUser))
	router.HandleFunc("/users/update", middleware.AuthMiddleware(userHandler.UpdateUser))
	router.HandleFunc("/users/delete", middleware.AuthMiddleware(userHandler.DeleteUser))
	router.HandleFunc("/users/username", middleware.AuthMiddleware(userHandler.FindUserByUsername))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	return &Server{
		router: router,
		port:   port,
	}
}

// Start begins listening for HTTP requests on the configured port.
func (s *Server) Start() error {
	log.Println("Server started on port", s.port)
	return http.ListenAndServe(":"+s.port, s.router)
}
