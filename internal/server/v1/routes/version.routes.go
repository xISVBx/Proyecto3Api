package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func VersionRoutes(r *gin.RouterGroup, c controllers.VersionController) {

	r.GET("version", c.GetVersion)

}
