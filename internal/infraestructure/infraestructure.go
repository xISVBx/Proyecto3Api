package infraestructure

import (
	"gorm.io/gorm"

	"col-moda/internal/infraestructure/repositories"
)

// Estructura que contiene todos los repositorios
type Infraestructure struct {
	UnitOfWork    *repositories.UnitOfWork
}

// Función para inicializar todos los repositorios
func InitInfraestructure(db *gorm.DB) *Infraestructure {
	return &Infraestructure{
		UnitOfWork: repositories.NewUnitOfWork(db),
	}
}
