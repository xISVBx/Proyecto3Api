package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func DepartmentRoutes(r *gin.RouterGroup, c controllers.DepartmentController) {

	r.GET("departments", c.FindAllDepartmentsByFilters)

}
