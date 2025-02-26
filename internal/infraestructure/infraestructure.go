package infraestructure

import (
	"gorm.io/gorm"

	"col-moda/internal/infraestructure/repositories"
)

// Estructura que contiene todos los repositorios
type Infraestructure struct {
	UserRepo *repositories.UserRepository
	RoleRepo *repositories.RoleRepository
}

// Función para inicializar todos los repositorios
func InitInfraestructure(db *gorm.DB) *Infraestructure {
	return &Infraestructure{
		UserRepo: repositories.NewUserRepository(db),
		RoleRepo: repositories.NewRoleRepository(db),
	}
}
