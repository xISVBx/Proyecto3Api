package controllers

import (
	"col-moda/internal/configuration"
	"col-moda/internal/domain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VersionController struct{

}

func NewVersionController() *VersionController{
	return &VersionController{}
}

// GetVersion obtiene la versión actual de la aplicación.
// @Summary Obtener versión
// @Description Retorna la versión de la API configurada en el sistema.
// @Tags Version
// @Produce json
// @Success 200 {object} models.AppResponse{data=string}
// @Router /api/v1/version [get]
func (ct VersionController) GetVersion(c *gin.Context){
	c.JSON(http.StatusOK, models.ResSuccess(configuration.AppConfiguration.Version))
}