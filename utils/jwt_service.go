package utils

import (
	"github.com/amardikamahdi/HL-Go/dtos"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateTokenJwt(email string) (dtos.AuthResponse, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		email: email,
	})

	mySigningKey := []byte("s3cr3t")

	signedString, err := token.SignedString(mySigningKey)
	if err != nil {
		return dtos.AuthResponse{}, err
	}

	return dtos.AuthResponse{Token: signedString}, nil
}
