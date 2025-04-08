// Package utils provides utility functions for handling HTTP responses, including JSON encoding and error handling.
package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/JonasBorgesLM/go_server/internal/dto"
)

// RespondWithError sends an HTTP response with a JSON-encoded error message, including the error message and field.
func RespondWithError(w http.ResponseWriter, statusCode int, message string, field string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(dto.ErrorResponse{
		Message: message,
		Field:   field,
	}); err != nil {
		log.Printf("failed to encode error response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// RespondWithJSON sends an HTTP response with a JSON-encoded payload and the specified status code.
func RespondWithJSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("failed to encode JSON response: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
