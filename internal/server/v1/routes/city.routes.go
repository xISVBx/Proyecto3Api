package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CityRoutes(r *gin.RouterGroup, c controllers.CityController) {

	r.GET("cities", c.FindCitiesByFilters)

}
