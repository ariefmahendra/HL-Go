package main

import (
	"github.com/amardikamahdi/HL-Go/configs"
	"github.com/amardikamahdi/HL-Go/routes"
	"github.com/amardikamahdi/HL-Go/structs"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	// DB configurations
	db := configs.ConnectDB()
	err := db.AutoMigrate(&structs.User{})
	if err != nil {
		log.Println(err)
		return
	}

	// Fiber initialization
	app := fiber.New()

	// Define routes
	routes.UserRoute(app)

	// Starting the runner
	err = app.Listen(":3000")
	if err != nil {
		log.Println(err)
		return
	}
}
