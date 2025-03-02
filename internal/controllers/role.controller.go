package controllers

import (
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"fmt"
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

// FindAllRoles Obtener todos los roles de la aplicacion
// @Summary Obtener todos los roles
// @Description Retorna un listado de los roles de la aplicacion.
// @Tags Roles
// @Accept json
// @Produce json
// @Success 201 {object} models.AppResponse{data=[]entities.Role}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 409 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/roles [get]
func (ct RoleController) FindAllRoles(c *gin.Context) {

	fmt.Println("entro")

	roles, aErr := ct.roleS.FindAllRoles()

	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(roles)
	c.JSON(http.StatusOK, response)
}
