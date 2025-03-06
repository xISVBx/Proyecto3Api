package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type CompanyService struct {
	repo *repositories.CompanyRepository
}

func NewCompanyService(repo *repositories.CompanyRepository) *CompanyService {
	return &CompanyService{repo: repo}
}

// FindCompaniesByFilters busca empresas aplicando filtros opcionales
func (r *CompanyService) FindCompaniesByFilters(dto dtos.CompanyRequestDto) ([]dtos.CompanyResponseDto, *models.AppError) {
	var companies, err = r.repo.FindCompaniesByFilters(dto)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	companiesResponse := dtos.CompaniesResponseDtoFromEntities(companies)

	return companiesResponse, nil
}
