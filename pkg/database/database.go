// Package database handles PostgreSQL connections and operations.
package database

import (
	"fmt"

	"github.com/JonasBorgesLM/go_server/internal/config"
	"github.com/jmoiron/sqlx"

	// PostgreSQL driver for database/sql.
	_ "github.com/lib/pq"
)

// Connect establishes a PostgreSQL connection using environment variables.
func Connect() (*sqlx.DB, error) {
	dbHost, err := config.GetEnv("DB_HOST")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DB_HOST: %w", err)
	}

	dbPort, err := config.GetEnv("DB_PORT")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DB_PORT: %w", err)
	}

	dbUsername, err := config.GetEnv("DB_USERNAME")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DB_USERNAME: %w", err)
	}

	dbPassword, err := config.GetEnv("DB_PASSWORD")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DB_PASSWORD: %w", err)
	}

	dbName, err := config.GetEnv("DB_DATABASE")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve DB_DATABASE: %w", err)
	}

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUsername,
		dbPassword,
		dbName,
	)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping the database: %w", err)
	}

	return db, nil
}
