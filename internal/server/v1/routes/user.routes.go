package v1_routes

import (
	"col-moda/internal/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(r *gin.RouterGroup, c controllers.UserController) {

	r.POST("login", c.Login)
	r.POST("register", c.Register)

}
