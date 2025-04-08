// Package handler provides HTTP handlers for login operations.
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/JonasBorgesLM/go_server/internal/dto"
	"github.com/JonasBorgesLM/go_server/pkg/utils"
)

// Login handles user authentication and returns a JWT token if successful.
func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var request dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to decode request body", "body")
		return
	}

	if err := request.Validate(); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input data", "validation")
		return
	}

	token, err := h.userUseCase.Execute(request.Email, request.Password)
	if err != nil {
		utils.RespondWithError(w, http.StatusUnauthorized, "Invalid credentials", "credentials")
		return
	}

	response := map[string]string{"token": token}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Failed to encode response", "response")
		return
	}
}
