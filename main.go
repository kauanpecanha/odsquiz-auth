package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")

	app := fiber.New(fiber.Config{
		AppName: "ODS Quiz Auth Microservice",
	})

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	log.Fatal(app.Listen(":" + port))
}
