package dtos

import (
	"github.com/google/uuid"
)

type UserCompanyResponseDto struct {
	ID          uuid.UUID `json:"id"`
	CompanyID   uuid.UUID `json:"company_id"`
	CompanyName string    `json:"company_name"`
	RoleID      uuid.UUID `json:"role_id"`
	Name        string    `json:"name"`
	LastName    string    `json:"las_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
}
