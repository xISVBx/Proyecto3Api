package controllers

import "col-moda/internal/services"

// Estructura que contiene todos los controladores
type Controllers struct {
	VersionController  *VersionController
	UserController     *UserController
	RoleController     *RoleController
	CategoryController *CategoryController
	ProductController  *ProductController
}

// Función para inicializar todos los controladores
func InitControllers(services *services.Services) *Controllers {
	return &Controllers{
		VersionController:  NewVersionController(),
		UserController:     NewUserController(services.UserService),
		RoleController:     NewRoleController(services.RoleService),
		CategoryController: NewCategoryController(services.CategorieService),
		ProductController:  NewProductController(services.ProductService),
	}
}
