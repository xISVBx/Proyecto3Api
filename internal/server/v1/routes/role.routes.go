package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RoleRoutes(r *gin.RouterGroup, c controllers.RoleController) {

	r.GET("roles", c.FindAllRoles)

}
