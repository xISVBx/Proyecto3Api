package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type UserResponseDto struct {
	ID       uuid.UUID `json:"id"`
	RoleID   uuid.UUID `json:"role_id"`
	Name     string    `json:"name"`
	LastName string    `json:"last_name"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
}

func UserResponseDtoFromEntity(user entities.User) UserResponseDto {
	return UserResponseDto{
		ID:     user.ID,
		RoleID: user.RoleID,
		Name:   user.Name,
		LastName: user.LastName,
		Email:  user.Email,
		Phone: user.Phone,
	}
}
