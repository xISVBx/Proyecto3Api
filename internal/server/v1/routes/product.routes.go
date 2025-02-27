package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.RouterGroup, c controllers.ProductController) {

	r.GET("products", c.FindProductsByFilters)

}
