// Package handler provides HTTP handlers for user operations.
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/JonasBorgesLM/go_server/internal/dto"
	"github.com/JonasBorgesLM/go_server/internal/formatter"
	"github.com/JonasBorgesLM/go_server/internal/usecase"
	"github.com/JonasBorgesLM/go_server/pkg/utils"
)

// UserHandler provides HTTP handlers for user-related operations.
type UserHandler struct {
	userUseCase *usecase.UserUseCase
}

// NewUserHandler creates a new instance of UserHandler with the provided UserUseCase.
func NewUserHandler(userUseCase *usecase.UserUseCase) *UserHandler {
	return &UserHandler{userUseCase}
}

// ListUsers retrieves and returns a list of all users.
func (h *UserHandler) ListUsers(w http.ResponseWriter, _ *http.Request) {
	users, err := h.userUseCase.ListUsers()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error listing users", "list_users")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(users); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error encoding response", "response")
		return
	}
}

// CreateUser handles the creation of a new user.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var request dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to decode request body", "body")
		return
	}

	if err := request.Validate(); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input data", "validation")
		return
	}

	user, err := formatter.CreateUserRequestToDomain(request)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error hashing password", "hashing")
		return
	}

	if err := h.userUseCase.CreateUser(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error creating user", "create_user")
		return
	}

	userResponse := formatter.UserToCreateUserResponse(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error encoding response", "response")
		return
	}
}

// UpdateUser handles the update of an existing user by ID.
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "User ID is required", "id")
		return
	}

	var request dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Failed to decode request body", "body")
		return
	}

	if err := request.Validate(); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid input data", "validation")
		return
	}

	if err := h.userUseCase.UpdateUser(id, &request); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error updating user", "update_user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(request); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error encoding response", "response")
		return
	}
}

// DeleteUser handles the deletion of a user by ID.
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "User ID is required", "id")
		return
	}

	if err := h.userUseCase.DeleteUser(id); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error deleting user", "delete_user")
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "User deleted successfully"})
}

// FindUserByUsername retrieves a user by their username.
func (h *UserHandler) FindUserByUsername(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		utils.RespondWithError(w, http.StatusBadRequest, "Username is required", "username")
		return
	}

	user, err := h.userUseCase.FindUserByUsername(username)
	if err != nil {
		utils.RespondWithError(w, http.StatusNotFound, "User not found", "find_user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, "Error encoding response", "response")
		return
	}
}
