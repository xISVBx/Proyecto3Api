package controllers

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CityController struct {
	cityS *services.CityService
}

func NewCityController(s *services.CityService) *CityController {
	return &CityController{
		cityS: s,
	}
}

// FindCitiesByFilters Obtener todas las ciudades de la aplicación
// @Summary Obtener todas las ciudades de la aplicación
// @Description Retorna un listado de las ciudades filtradas por departamento o nombre de ciudad.
// @Tags Cities
// @Accept json
// @Produce json
// @Param department_id query int false "ID del departamento (opcional)"
// @Param city_name query string false "Nombre de la ciudad (opcional)"
// @Param city_id query int false "ID de la ciudad (opcional)"
// @Success 200 {object} models.AppResponse{data=[]dtos.CityResponseDto}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 409 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/cities [get]
func (ct CityController) FindCitiesByFilters(c *gin.Context) {

	var dto dtos.CityRequestDto
	err := c.ShouldBindQuery(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResError(err.Error()))
		return
	}

	cities, aErr := ct.cityS.FindCitiesByFilters(dto)

	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(cities)
	c.JSON(http.StatusOK, response)
}
