package routes

import "github.com/gofiber/fiber/v3"

func Setup(app *fiber.App) {
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("ok")
	})
}