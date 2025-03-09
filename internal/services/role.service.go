package services

import (
	"col-moda/internal/domain/entities"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type RoleService struct {
	uow *repositories.UnitOfWork
}

func NewRoleService(uow *repositories.UnitOfWork) *RoleService {
	return &RoleService{
		uow: uow,
	}
}

func (s RoleService) FindAllRoles() ([]entities.Role, *models.AppError) {

	s.uow.Begin()

	defer s.uow.Commit()
	
	roles, err := s.uow.RoleRepo.FindAllRoles()

	if err != nil {
		return nil, models.NewServerError(err)
	}

	return roles, nil
}
