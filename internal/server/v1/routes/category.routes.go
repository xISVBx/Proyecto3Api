package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.RouterGroup, c controllers.CategoryController) {

	r.GET("categories", c.FindAllCategories)

}
