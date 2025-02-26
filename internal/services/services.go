package services

import (
	"col-moda/internal/infraestructure"
)

// Estructura que contiene todos los servicios
type Services struct {
	UserService *UserService
	RoleService *RoleService
}

// Función para inicializar todos los servicios
func InitServices(repos *infraestructure.Infraestructure) *Services {
	return &Services{
		UserService: NewUserService(repos.UserRepo, repos.RoleRepo),
		RoleService: NewRoleService(repos.RoleRepo),
	}
}
