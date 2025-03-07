package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type DepartmentService struct {
	uow *repositories.UnitOfWork
}

func NewDepartmentService(uow *repositories.UnitOfWork) *DepartmentService {
	return &DepartmentService{
		uow: uow,
	}
}

func (s DepartmentService) FindAllDepartmentsByFilters(dto dtos.DepartmentRequestDto) ([]dtos.DepartmentResponseDto, *models.AppError) {
	departments, err := s.uow.DepartmentRepo.FindAllDepartmentsByFilters(dto)
	if err != nil {
		return nil, models.NewServerError(err)
	}

	departmentsResponse := dtos.DepartmentsResponseFromEntitie(departments)

	return departmentsResponse, nil
}
