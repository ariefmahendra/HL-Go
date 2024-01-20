package repositories

import (
	"github.com/amardikamahdi/HL-Go/configs"
	"github.com/amardikamahdi/HL-Go/structs"
	"github.com/amardikamahdi/HL-Go/utils"
	"github.com/amardikamahdi/HL-Go/validator"
	"github.com/gofiber/fiber/v2"
)

func GetAllUsers(c *fiber.Ctx) error {
	var users []structs.User
	err := configs.ConnectDB().Find(&users).Error

	if err != nil {
		return c.SendStatus(500)
	}

	return c.JSON(users)
}

func GetUserById(c *fiber.Ctx) error {
	var user structs.User
	err := configs.ConnectDB().First(&user, c.Params("id")).Error

	if err != nil {
		return c.SendStatus(404)
	}

	return c.JSON(user)
}

func CreateUser(c *fiber.Ctx) error {
	var user structs.User

	if err := c.BodyParser(&user); err != nil {
		return c.SendStatus(422)
	}

	if err := validator.CreateUserValidation(user); err != nil {
		return c.Status(422).JSON(err)
	}

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return c.SendStatus(500)
	}

	user.Password = hashedPassword

	if createErr := configs.ConnectDB().Create(&user).Error; createErr != nil {
		return c.SendStatus(500)
	}

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user structs.User

	findErr := configs.ConnectDB().First(&user, c.Params("id")).Error

	if findErr != nil {
		return c.SendStatus(404)
	}

	deleteErr := configs.ConnectDB().Delete(&structs.User{}, c.Params("id")).Error

	if deleteErr != nil {
		return c.SendStatus(500)
	}

	return c.SendStatus(204)
}
