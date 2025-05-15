package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	
	"draw/internal/config"
	"draw/internal/models"
)

// Initialize sets up the database connection and creates necessary tables
func Initialize(cfg *config.Config) (*sql.DB, error) {
	// Open database connection
	db, err := sql.Open("mysql", cfg.GetDSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// Check if connection is alive
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("database ping failed: %w", err)
	}

	log.Println("Connected to database successfully")

	// Create tables if they don't exist
	if err = createTables(db); err != nil {
		return nil, fmt.Errorf("failed to create tables: %w", err)
	}

	return db, nil
}

// createTables creates all necessary tables in the database
func createTables(db *sql.DB) error {
	// Create users table
	if err := models.CreateUsersTable(db); err != nil {
		return err
	}

	log.Println("Database tables created/verified successfully")
	return nil
}
