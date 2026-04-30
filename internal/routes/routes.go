package routes

import (
	"encoding/json"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
)

// setup function to config routes
func Setup(app *fiber.App) {
	
	// sample users to simulate database
	sampleUsers := []models.User{
		{
			Name:     "João Silva",
			Email:    "joao.silva@example.com",
			Password: "password123",
		},
		{
			Name:     "Maria Santos",
			Email:    "maria.santos@example.com",
			Password: "password456",
		},
		{
			Name:     "Pedro Oliveira",
			Email:    "pedro.oliveira@example.com",
			Password: "password789",
		},
	}

	// root route
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

	// get all users records route
	app.Get("/getAllUsers", func(c fiber.Ctx) error {
		return c.JSON(sampleUsers)
	})

	// get a single user by its id route
	app.Get("/getUserById/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		for _, user := range sampleUsers {
			if user.ID == id {
				return c.JSON(user)
			}
		}
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	})

	// route to create user
	app.Post("/createUser", func(c fiber.Ctx) error {
		var newUser models.User
		body := c.BodyRaw()
		if len(body) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "empty request body"})
		}
		if err := json.Unmarshal(body, &newUser); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		if newUser.ID == "" {
			newUser.ID = uuid.NewString()
		}

		sampleUsers = append(sampleUsers, newUser)
		return c.Status(fiber.StatusCreated).JSON(sampleUsers)
	})

	// update user attributes route
	app.Patch("/updateUser/:id", func(c fiber.Ctx) error {
		id := c.Params("id")
		var payload map[string]any
		body := c.BodyRaw()
		if len(body) == 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "empty request body"})
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request body"})
		}

		for idx, user := range sampleUsers {
			if user.ID != id {
				continue
			}

			if name, ok := payload["name"].(string); ok {
				sampleUsers[idx].Name = name
			}
			if email, ok := payload["email"].(string); ok {
				sampleUsers[idx].Email = email
			}
			if password, ok := payload["password"].(string); ok {
				sampleUsers[idx].Password = password
			}

			sampleUsers[idx].UpdatedAt = time.Now()
			return c.JSON(sampleUsers[idx])
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	})

	// delete user by id route
	app.Delete("/deleteUser/:id", func(c fiber.Ctx) error {
		id := c.Params("id")

		for idx, user := range sampleUsers {
			if user.ID == id {
				deleted := user
				sampleUsers = append(sampleUsers[:idx], sampleUsers[idx+1:]...)
				return c.JSON(fiber.Map{
					"message": "user deleted",
					"user":    deleted,
				})
			}
		}

		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "user not found"})
	})
}
