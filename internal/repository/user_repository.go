// Package repository provides implementations for accessing and manipulating data in the user database.
package repository

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JonasBorgesLM/go_server/internal/domain"
	"github.com/JonasBorgesLM/go_server/internal/dto"
	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

// NewUserRepository creates a new instance of userRepository with the provided database connection.
func NewUserRepository(db *sqlx.DB) domain.UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
	query := `INSERT INTO users (username, email, password) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, user.Username, user.Email, user.Password)
	if err != nil {
		log.Printf("failed to create user: %v", err)
		return fmt.Errorf("failed to create user: %v", err)
	}
	return nil
}

func (r *userRepository) FindByEmail(email string) (*domain.User, error) {
	var user domain.User
	query := `SELECT id, email, password FROM users WHERE email = $1`
	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user with email %s not found", email)
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("failed to find user by email: %v", err)
		return nil, fmt.Errorf("failed to find user by email: %v", err)
	}
	return &user, nil
}

func (r *userRepository) FindByUsername(username string) (*dto.UserResponse, error) {
	var user dto.UserResponse
	query := `SELECT id, username, email FROM users WHERE username = $1`
	err := r.db.QueryRow(query, username).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user with username %s not found", username)
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("failed to find user by username: %v", err)
		return nil, fmt.Errorf("failed to find user by username: %v", err)
	}
	return &user, nil
}

func (r *userRepository) FindByID(id string) (*dto.UserResponse, error) {
	var user dto.UserResponse
	query := `SELECT id, username, email FROM users WHERE id = $1`
	err := r.db.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("user with id %s not found", id)
			return nil, fmt.Errorf("user not found")
		}
		log.Printf("failed to find user by id: %v", err)
		return nil, fmt.Errorf("failed to find user by id: %v", err)
	}
	return &user, nil
}

func (r *userRepository) List() (*[]dto.UserResponse, error) {
	query := `SELECT id, username, email FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		log.Printf("failed to execute query: %v", err)
		return nil, fmt.Errorf("failed to execute query: %v", err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("failed to close rows: %v", err)
		}
	}()

	var users []dto.UserResponse
	for rows.Next() {
		var user dto.UserResponse
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			log.Printf("failed to scan user data: %v", err)
			return nil, fmt.Errorf("failed to scan user data: %v", err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		log.Printf("error after row iteration: %v", err)
		return nil, fmt.Errorf("error after row iteration: %v", err)
	}

	return &users, nil
}

func (r *userRepository) Update(id string, user *dto.UpdateUserRequest) error {
	query := `UPDATE users SET username = $1, email = $2 WHERE id = $3`
	_, err := r.db.Exec(query, user.Username, user.Email, id)
	if err != nil {
		log.Printf("failed to update user: %v", err)
		return fmt.Errorf("failed to update user: %v", err)
	}
	return nil
}

func (r *userRepository) Delete(id string) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		log.Printf("failed to delete user: %v", err)
		return fmt.Errorf("failed to delete user: %v", err)
	}
	return nil
}
