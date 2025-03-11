package jwt

import (
	"col-moda/internal/configuration"
	"col-moda/internal/domain/dtos"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func CreateUserToken(user dtos.UserResponseDto) (*string, error) {
	// Define claims (payload)
	claims := jwt.MapClaims{
		"id":        user.ID,
		"role_id":   user.RoleID,
		"role":      "Usuario",
		"name":      user.Name,
		"last_name": user.LastName,
		"email":     user.Email,
		"phone":     user.Phone,
		"exp":       time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token, err := createToken(claims)

	if err != nil {
		return nil, err
	}

	return token, nil

}

func CreateUserCompanyToken(company dtos.UserCompanyResponseDto) (*string, error) {

	claims := jwt.MapClaims{
		"id":           company.ID,
		"company_id":   company.CompanyID,
		"company_name": company.CompanyName,
		"role_id":      company.RoleID,
		"role":         "Company",
		"name":         company.Name,
		"last_name":    company.LastName,
		"email":        company.Email,
		"phone":        company.Phone,
		"exp":          time.Now().Add(time.Hour * 24).Unix(), // Expires in 24 hours
	}

	token, err := createToken(claims)

	if err != nil {
		return nil, err
	}

	return token, nil
}

func createToken(claims jwt.MapClaims) (*string, error) {
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign token with secret key
	signedToken, err := token.SignedString(configuration.AppConfiguration.Jwt)
	if err != nil {
		return nil, err
	}

	return &signedToken, nil
}

func VerifyToken(tokenString string) (jwt.MapClaims, error) {
	// Parsear el token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Verificar que el método de firma sea HMAC
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("método de firma inválido")
		}
		return configuration.AppConfiguration.Jwt, nil // Usar la clave secreta
	})

	if err != nil || !token.Valid {
		return nil, errors.New("token inválido o expirado")
	}

	// Extraer claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("no se pudieron obtener los claims del token")
	}

	return claims, nil
}
