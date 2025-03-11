package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type RegisterResponseDto struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
}


func RegisterResponseDtoFromEntitie(user entities.User) RegisterResponseDto {
	return RegisterResponseDto{
		ID:       user.ID,
		Name:     user.Name,
		LastName: user.LastName,
		Email:    user.Email,
	}
}
