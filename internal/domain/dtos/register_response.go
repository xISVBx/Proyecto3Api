package dtos

import "col-moda/internal/domain/entities"

type RegisterResponseDto struct {
	Name     string
	LastName string
	Email    string
}

func RegisterResponseDtoFromEntitie(user entities.User) RegisterResponseDto {
	return RegisterResponseDto{
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
}
