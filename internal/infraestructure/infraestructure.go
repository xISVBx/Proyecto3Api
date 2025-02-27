package infraestructure

import (
	"gorm.io/gorm"

	"col-moda/internal/infraestructure/repositories"
)

// Estructura que contiene todos los repositorios
type Infraestructure struct {
	UserRepo     *repositories.UserRepository
	RoleRepo     *repositories.RoleRepository
	CategoryRepo *repositories.CategoryRepository
	ProductRepo  *repositories.ProductRepository
}

// Función para inicializar todos los repositorios
func InitInfraestructure(db *gorm.DB) *Infraestructure {
	return &Infraestructure{
		UserRepo:     repositories.NewUserRepository(db),
		RoleRepo:     repositories.NewRoleRepository(db),
		CategoryRepo: repositories.NewCategoriesRepository(db),
		ProductRepo: repositories.NewProductRepository(db),
	}
}
