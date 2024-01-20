package routes

import (
	"github.com/amardikamahdi/HL-Go/repositories"
	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	// GET /users
	app.Get("/users", getAllUsers)

	// GET /users/:id
	app.Get("/users/:id", getUserById)

	// POST /users
	app.Post("/users", createUser)

	// DELETE /users/:id
	app.Delete("/users/:id", deleteUser)
}

func getAllUsers(c *fiber.Ctx) error {
	return repositories.GetAllUsers(c)
}

func getUserById(c *fiber.Ctx) error {
	return repositories.GetUserById(c)
}

func createUser(c *fiber.Ctx) error {
	return repositories.CreateUser(c)
}

func deleteUser(c *fiber.Ctx) error {
	return repositories.DeleteUser(c)
}
