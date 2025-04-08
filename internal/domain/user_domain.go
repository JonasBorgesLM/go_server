// Package domain defines the core entities and interfaces related to user management.
package domain

import "github.com/JonasBorgesLM/go_server/internal/dto"

// User represents a user entity with its essential attributes.
type User struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
}

// UserRepository defines the contract for interacting with user data storage.
type UserRepository interface {
	FindByEmail(email string) (*User, error)
	FindByID(id string) (*dto.UserResponse, error)
	FindByUsername(username string) (*dto.UserResponse, error)
	List() (*[]dto.UserResponse, error)
	Create(user *User) error
	Update(id string, user *dto.UpdateUserRequest) error
	Delete(id string) error
}
