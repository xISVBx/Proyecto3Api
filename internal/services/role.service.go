package services

import (
	"col-moda/internal/domain/entities"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type RoleService struct {
	roleR *repositories.RoleRepository
}

func NewRoleService(r *repositories.RoleRepository) *RoleService {
	return &RoleService{
		roleR: r,
	}
}

func (s RoleService) FindAllRoles() ([]entities.Role, *models.AppError) {
	roles, err := s.roleR.FindAllRoles()

	if err != nil {
		return nil, models.NewServerError(err)
	}

	return roles, nil
}
