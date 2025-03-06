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
