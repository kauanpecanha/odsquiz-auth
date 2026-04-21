package main

import (
	"log"

	"github.com/gofiber/fiber/v3"

	"github.com/kauanpecanha/odsquiz-auth/pkg/config"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
	"github.com/kauanpecanha/odsquiz-auth/api/routes"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewPostgresConnection(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pgDb, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	defer pgDb.Close()

	app := fiber.New(fiber.Config{
		AppName: "ODS Quiz Auth Microservice",
	})

	routes.Setup(app)

	log.Fatal(app.Listen(":" + cfg.Port))
}