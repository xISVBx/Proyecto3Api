package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type RegisterResponseDto struct {
	ID       uuid.UUID
	Name     string
	LastName string
	Email    string
}

func RegisterResponseDtoFromEntitie(user entities.User) RegisterResponseDto {
	return RegisterResponseDto{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
}
