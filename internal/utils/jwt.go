package utils

import (
	"col-moda/internal/configuration"
	"col-moda/internal/domain/dtos"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateUserToken(user dtos.UserGetDto) (*string, error) {
	// Define claims (payload)
	claims := jwt.MapClaims{
		"name":  user.Name,
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, err := token.SignedString(configuration.AppConfiguration.Jwt)
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}
