package server

import (
	//middleware "col-moda/internal/server/middlewares"
	v1_routes "col-moda/internal/server/v1/routes"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (s *Server) RegisterRoutes() (http.Handler, error) {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Add your frontend URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type"},
		AllowCredentials: true, // Enable cookies/auth
	}))

	//protected := r.Group("/api/v1")
	//protected.Use(middleware.AuthMiddleware())

	//📄Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register API routes
	v1 := r.Group("/api/v1")
	{
		v1_routes.UserRoutes(v1, *s.Controllers.UserController)
		v1_routes.VersionRoutes(v1, *s.Controllers.VersionController)
		v1_routes.RoleRoutes(v1, *s.Controllers.RoleController)
		v1_routes.CategoryRoutes(v1, *s.Controllers.CategoryController)
		v1_routes.ProductRoutes(v1, *s.Controllers.ProductController)
		v1_routes.CityRoutes(v1, *s.Controllers.CityController)
		v1_routes.DepartmentRoutes(v1, *s.Controllers.DepartmentController)
	}

	

	return r, nil
}
