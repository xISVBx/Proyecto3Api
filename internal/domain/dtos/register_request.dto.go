package dtos

import (
	"col-moda/internal/domain/entities"
	"github.com/google/uuid"
)

type RegisterRequest struct{
	Name string
	LastName string
	Email string
	Password string
	CityId int
	Address string
	Phone string
}

func RegisterRequestToUserEntity(registerRequest *RegisterRequest, roleId uuid.UUID) entities.User{
	return entities.User{
		ID: uuid.New(),
		Name: registerRequest.Name,
		LastName: registerRequest.LastName,
		Email: registerRequest.Email,
		Password: registerRequest.Password,
		CityID: registerRequest.CityId,
		Address: registerRequest.Address,
		Phone: registerRequest.Phone,
		RoleID: roleId,
	}
}