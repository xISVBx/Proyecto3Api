package services

import (
	"col-moda/internal/infraestructure"
)

// Estructura que contiene todos los servicios
type Services struct {
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
		DepartmentService: NewDepartmentService(repos.DepartmentRepo), UserService: NewUserService(repos.UserRepo, repos.RoleRepo),
		RoleService:      NewRoleService(repos.RoleRepo),
		CategorieService: NewCategorieService(repos.CategoryRepo),
		ProductService:   NewProductService(repos.ProductRepo),
		CityService:      NewCityService(repos.CityRepo),
	}
}
