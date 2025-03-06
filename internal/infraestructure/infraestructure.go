package infraestructure

import (
	"gorm.io/gorm"

	"col-moda/internal/infraestructure/repositories"
)

// Estructura que contiene todos los repositorios
type Infraestructure struct {
	DepartmentRepo *repositories.DepartmentRepository
	UserRepo       *repositories.UserRepository
	RoleRepo       *repositories.RoleRepository
	CategoryRepo   *repositories.CategoryRepository
	ProductRepo    *repositories.ProductRepository
	CityRepo       *repositories.CityRepository
	CompanyRepo   *repositories.CompanyRepository
}

// Función para inicializar todos los repositorios
func InitInfraestructure(db *gorm.DB) *Infraestructure {
	return &Infraestructure{
		DepartmentRepo: repositories.NewDepartmentRepository(db), UserRepo: repositories.NewUserRepository(db),
		RoleRepo:     repositories.NewRoleRepository(db),
		CategoryRepo: repositories.NewCategoriesRepository(db),
		ProductRepo:  repositories.NewProductRepository(db),
		CityRepo:     repositories.NewCityRepository(db),
		CompanyRepo: repositories.NewCompanyRepository(db),
	}
}
