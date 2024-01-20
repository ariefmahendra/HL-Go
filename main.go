package main

import (
	"github.com/amardikamahdi/HL-Go/configs"
	"github.com/amardikamahdi/HL-Go/routes"
	"github.com/amardikamahdi/HL-Go/structs"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// DB configurations
	db := configs.ConnectDB()
	db.AutoMigrate(&structs.User{})

	// Fiber initialization
	app := fiber.New()

	// Define routes
	routes.UserRoute(app)

	// Starting the runner
	app.Listen(":3000")
}
