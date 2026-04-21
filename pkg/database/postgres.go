// Package database provides functions for database connections and operations.
package database

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
)

// NewPostgresConnection establishes a connection to the PostgreSQL database using GORM.
func NewPostgresConnection(cfg *config.Config) (*gorm.DB, error) {
	// Build the database connection string from config
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
		cfg.DBSSLMode,
	)

	// Open the database connection
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Log successful connection
	log.Println("Successfull database connection")

	return db, nil
}
