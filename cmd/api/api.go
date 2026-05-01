// Package main is the entry point for the ODS Quiz Auth API microservice.
package main

import (
	"log"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"

	"github.com/kauanpecanha/odsquiz-auth/internal/routes"
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
	_, err = database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}
	
	// Create new Fiber app instance
	app := fiber.New(fiber.Config{
		AppName: "ODS Quiz Auth Microservice",
	})

	app.Use(cors.New())

	// Set up API routes
	routes.Setup(app)

	app.Use(func(c fiber.Ctx) error {
		// 404 "Not Found"
		return c.SendStatus(fiber.StatusNotFound)
	})

	// Start the server on the configured port
	log.Fatal(app.Listen(":" + cfg.Port))
}
