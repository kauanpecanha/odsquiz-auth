package routes

import "github.com/gofiber/fiber/v3"

func Setup(app *fiber.App) {
	// app health verification handler
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})

	// 404 handler
	app.Use(func(c fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})
}