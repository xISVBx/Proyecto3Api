package controllers

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DepartmentController struct {
	service *services.DepartmentService
}

func NewDepartmentController(service *services.DepartmentService) *DepartmentController {
	return &DepartmentController{service: service}
}

// FindCitiesByFilters Obtener todas las ciudades de la aplicación
// @Summary Obtener todos los departamentos de la aplicación
// @Description Retorna un listado de las ciudades filtradas por departamento o nombre de ciudad.
// @Tags Departments
// @Accept json
// @Produce json
// @Param department_id query int false "ID del departamento (opcional)"
// @Param department_name query string false "Nombre de la ciudad (opcional)"
// @Success 200 {object} models.AppResponse{data=[]dtos.DepartmentRequestDto}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 409 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/departments [get]
func (ct DepartmentController) FindAllDepartmentsByFilters(c *gin.Context) {

	var dto dtos.DepartmentRequestDto
	err := c.ShouldBindQuery(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResError(err.Error()))
		return
	}

	departments, aErr := ct.service.FindAllDepartmentsByFilters(dto)
	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(departments)
	c.JSON(http.StatusOK, response)
}
