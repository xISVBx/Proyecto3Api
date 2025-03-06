package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type RegisterRequestDto struct {
	Name     string `json:"name" binding:"required"`
	LastName string `json:"last_name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	CityId   int    `json:"city_id"`
	Address  string `json:"address"`
	Phone    string `json:"phone"`
}

func RegisterRequestDtoToEntitie(registerRequest *RegisterRequestDto, roleId uuid.UUID, hashedPassword string) entities.User {
	return entities.User{
		ID:             uuid.New(),
		Name:           registerRequest.Name,
		LastName:       registerRequest.LastName,
		Email:          registerRequest.Email,
		HashedPassword: hashedPassword,
		CityID:         registerRequest.CityId,
		Address:        registerRequest.Address,
		Phone:          registerRequest.Phone,
		RoleID:         roleId,
	}
}
