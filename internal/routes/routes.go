package routes

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kauanpecanha/odsquiz-auth/internal/models"
	"github.com/kauanpecanha/odsquiz-auth/pkg/database"
)

// setup function to config routes
func Setup(app *fiber.App) {

	// root route
	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello, world!")
	})

		// route to create user
	app.Post("/createUser", func(c fiber.Ctx) error {
		user := new(models.User)
		if user.ID == "" {
			user.ID = uuid.NewString()
		}
		if err := c.Bind().Body(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		if result := database.DB.Db.Create(&user); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
		}

		return c.Status(fiber.StatusOK).JSON(user)
	})

	// get all users records route
	app.Get("/getAllUsers", func(c fiber.Ctx) error {
		users := []models.User{}

		if result := database.DB.Db.Find(&users); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
		}

		return c.Status(fiber.StatusOK).JSON(users)
	})

	// get a single user by its id route
	app.Get("/getUserById/:id", func(c fiber.Ctx) error {
		user := models.User{}

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("invalid id")
		}

		if result := database.DB.Db.First(&user, id); result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(result.Error.Error())
		}

		return c.Status(fiber.StatusOK).JSON(user)
	})

	// update user attributes route
	app.Patch("/updateUser/:id", func(c fiber.Ctx) error {
		user := new(models.User)
		if err := c.Bind().Body(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(err.Error())
		}

		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("invalid id")
		}

		if result := database.DB.Db.Model(&models.User{}).Where("id = ?", id).Updates(user); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
		}

		return c.Status(fiber.StatusOK).JSON("updated")
	})

	// delete user by id route
	app.Delete("/deleteUser/:id", func(c fiber.Ctx) error {
		id, err := uuid.Parse(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON("invalid id")
		}

		if result := database.DB.Db.Delete(&models.User{}, id); result.Error != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(result.Error.Error())
		}

		return c.Status(fiber.StatusOK).JSON("deleted")
	})
}
