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

	// Login
	app.Post("/auth/login", login)

	// register
	app.Post("/auth/register", register)
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

func login(c *fiber.Ctx) error {
	return repositories.Login(c)
}

func register(c *fiber.Ctx) error {
	return repositories.Register(c)
}
