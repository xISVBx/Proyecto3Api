package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type DepartmentService struct {
	repo *repositories.DepartmentRepository
}

func NewDepartmentService(repo *repositories.DepartmentRepository) *DepartmentService {
	return &DepartmentService{repo: repo}
}

func (s DepartmentService) FindAllDepartmentsByFilters(dto dtos.DepartmentRequest) ([]dtos.DepartmentResponse, *models.AppError) {
	departments, err := s.repo.FindAllDepartmentsByFilters(dto)
	if err != nil {
		return nil, models.NewServerError(err)
	}

	departmentsResponse := dtos.DepartmentsResponseFromEntitie(departments)

	return departmentsResponse, nil
}
