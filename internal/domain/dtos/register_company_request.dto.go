package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type RegisterCompanyRequestDto struct {
	Name        string `json:"name" binding:"required"`           // Obligatorio
	LastName    string `json:"last_name" binding:"required"`      // Obligatorio
	Email       string `json:"email" binding:"required,email"`    // Obligatorio y debe ser un email válido
	Password    string `json:"password" binding:"required,min=8"` // Mínimo 8 caracteres
	CityId      int    `json:"city_id" binding:"gt=0"`            // Debe ser mayor que 0
	Address     string `json:"address" binding:"required"`        // Obligatorio
	Phone       string `json:"phone" binding:"omitempty,numeric"` // Opcional pero debe ser numérico
	CompanyName string `json:"company_name" binding:"required"`   // Obligatorio
	Description string `json:"description"`
} 

func RegisterCompanyRequestDtoToEntitie(registerCompanyRequest RegisterCompanyRequestDto, roleID uuid.UUID) (entities.User, entities.Company, entities.UserCompany) {

	userID := uuid.New() 
	companyID := uuid.New()

	return entities.User{
			ID:             userID,
			Name:           registerCompanyRequest.Name,
			LastName:       registerCompanyRequest.LastName,
			Email:          registerCompanyRequest.Email,
			HashedPassword: registerCompanyRequest.Password,
			CityID:         registerCompanyRequest.CityId,
			Address:        registerCompanyRequest.Address,
			Phone:          registerCompanyRequest.Phone,
			RoleID:         roleID,
		}, entities.Company{
			ID:          companyID,
			CompanyName:        registerCompanyRequest.CompanyName,
			Description: registerCompanyRequest.Description,
		}, entities.UserCompany{
			UserID:    userID,
			CompanyID: companyID,
		}
}
