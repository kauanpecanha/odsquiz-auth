// Package main is the entry point for the ODS Quiz Auth API microservice.
package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/kauanpecanha/odsquiz-auth/api/routes"
	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
)

// main initializes and starts the API server.
func main() {
	// Load configuration settings
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Establish connection to PostgreSQL database
	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	// Get underlying database connection
	pgDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer pgDb.Close()

	// Create new Fiber app instance
	app := fiber.New(fiber.Config{
		AppName: "ODS Quiz Auth Microservice",
	})

	// Set up API routes
	routes.Setup(app)

	// Start the server on the configured port
	log.Fatal(app.Listen(":" + cfg.Port))
}
