package repositories

import (
	"github.com/amardikamahdi/HL-Go/configs"
	"github.com/amardikamahdi/HL-Go/dtos"
	"github.com/amardikamahdi/HL-Go/structs"
	"github.com/amardikamahdi/HL-Go/utils"
	"github.com/amardikamahdi/HL-Go/validator"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Login(c *fiber.Ctx) error {
	var login dtos.AuthRequest

	err := c.BodyParser(&login)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	var user structs.User
	user.Email = login.Email
	err = configs.ConnectDB().First(&user).Error
	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
	if err != nil {
		return c.SendStatus(http.StatusUnauthorized)
	}

	authResponse, err := utils.GenerateTokenJwt(user.Email)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(authResponse)
}

func Register(c *fiber.Ctx) error {

	var userRequest structs.User

	if err := c.BodyParser(&userRequest); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	if err := validator.CreateUserValidation(userRequest); err != nil {
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	passwordHashed, err := utils.HashPassword(userRequest.Password)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	userRequest.Password = passwordHashed

	err = configs.ConnectDB().Create(&userRequest).Error
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.JSON(userRequest)
}
