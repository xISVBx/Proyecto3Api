package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type CompanyResponseDto struct {
	ID          uuid.UUID `form:"id,omitempty"`
	CompanyName string    `form:"company_name,omitempty"`
}

func CompaniesResponseDtoFromEntities(companies []entities.Company) []CompanyResponseDto {
	var companiesResponse []CompanyResponseDto
	for _, company := range companies {
		companiesResponse = append(companiesResponse, CompanyResponseDto{
			ID:          company.ID,
			CompanyName: company.CompanyName,
		})
	}
	return companiesResponse
}
