package validator

import (
	"github.com/amardikamahdi/HL-Go/structs"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

func CreateUserValidation(user structs.User) error {
	return validation.ValidateStruct(&user,
		validation.Field(&user.Name, validation.Required),
		validation.Field(&user.Email, validation.Required, is.Email),
		validation.Field(&user.Password, validation.Required),
	)
}
