package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type RegisterCompanyResponseDto struct {
	CompanyID   uuid.UUID `json:"company_id"`
	UserID      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	CompanyName string    `json:"company_name"`
	Email       string    `json:"email"`
}


func RegisterCompanyResponseDtoFromEntities(user entities.User, company entities.Company) RegisterCompanyResponseDto {

	return RegisterCompanyResponseDto{
		CompanyID:   company.ID,
		UserID:      user.ID,
		Name:        user.Name,
		CompanyName: company.CompanyName,
		Email:       user.Email,
	}
}
