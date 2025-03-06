package controllers

import "col-moda/internal/services"

// Estructura que contiene todos los controladores
type Controllers struct {
	CompanyController    *CompanyController
	DepartmentController *DepartmentController
	VersionController    *VersionController
	UserController       *UserController
	RoleController       *RoleController
	CategoryController   *CategoryController
	ProductController    *ProductController
	CityController       *CityController
}

// Función para inicializar todos los controladores
func InitControllers(services *services.Services) *Controllers {
	return &Controllers{
		DepartmentController: NewDepartmentController(services.DepartmentService), VersionController: NewVersionController(),
		UserController:     NewUserController(services.UserService),
		RoleController:     NewRoleController(services.RoleService),
		CategoryController: NewCategoryController(services.CategorieService),
		ProductController:  NewProductController(services.ProductService),
		CityController:     NewCityController(services.CityService),
		CompanyController: NewCompaniesController(services.CompanyService),
	}
}
