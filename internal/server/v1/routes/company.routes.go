package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func CompanyRoutes(r *gin.RouterGroup, c controllers.CompanyController) {

	r.GET("companies", c.FindAllCompaniesByFilters)

	r.POST("companies", c.RegisterCompany)
}
