package controllers

import (
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RoleController struct {
	roleS *services.RoleService
}

func NewRoleController(s *services.RoleService) *RoleController {
	return &RoleController{
		roleS: s,
	}
}

func (ct RoleController) FindAllRoles(c *gin.Context) {

	roles, aErr := ct.roleS.FindAllRoles()

	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(roles)
	c.JSON(http.StatusOK, response)
}
