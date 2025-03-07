package services

import (
	"col-moda/internal/infraestructure"
)

// Estructura que contiene todos los servicios
type Services struct {
	CompanyService    *CompanyService
	DepartmentService *DepartmentService
	UserService       *UserService
	RoleService       *RoleService
	CategorieService  *CategorieService
	ProductService    *ProductService
	CityService       *CityService
}

// Función para inicializar todos los servicios
func InitServices(repos *infraestructure.Infraestructure) *Services {
	return &Services{
		DepartmentService: NewDepartmentService(repos.UnitOfWork),
		UserService:       NewUserService(repos.UnitOfWork),
		RoleService:       NewRoleService(repos.UnitOfWork),
		CategorieService:  NewCategorieService(repos.UnitOfWork),
		ProductService:    NewProductService(repos.UnitOfWork),
		CityService:       NewCityService(repos.UnitOfWork),
		CompanyService:    NewCompanyService(repos.UnitOfWork),
	}
}
