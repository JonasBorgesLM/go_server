// Package dto contains data transfer objects used in the error response.
package dto

// ErrorResponse represents a standardized structure for error responses in the API.
type ErrorResponse struct {
	Message string `json:"message"`
	Field   string `json:"field,omitempty"`
}
