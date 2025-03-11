package controllers

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	service *services.CompanyService
}

func NewCompaniesController(service *services.CompanyService) *CompanyController {
	return &CompanyController{service: service}
}

// FindAllDepartmentsByFilters busca empresas aplicando filtros opcionales
// @Summary Busca empresas por filtros opcionales
// @Description Obtiene una lista de empresas aplicando filtros como ID o nombre
// @Tags Empresas
// @Accept  json
// @Produce  json
// @Param id query string false "ID de la empresa (UUID)"
// @Param company_name query string false "Nombre de la empresa"
// @Success 200 {object} models.AppResponse{data=[]dtos.CompanyResponseDto}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 500 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/companies [get]
func (ct CompanyController) FindAllCompaniesByFilters(c *gin.Context) {

	var dto dtos.CompanyRequestDto
	err := c.ShouldBindQuery(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResError(err.Error()))
		return
	}

	companies, aErr := ct.service.FindCompaniesByFilters(dto)
	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(companies)
	c.JSON(http.StatusOK, response)
}

// RegisterCompany registra una nueva empresa
// @Summary Registra una nueva empresa
// @Description Crea una nueva empresa con la información proporcionada
// @Tags Empresas
// @Accept  json
// @Produce  json
// @Param company body dtos.RegisterCompanyRequestDto true "Datos de la empresa"
// @Success 201 {object} models.AppResponse{data=dtos.RegisterCompanyResponseDto}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 500 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/companies [post]
func (ct CompanyController) RegisterCompany(c *gin.Context) {

	var dto dtos.RegisterCompanyRequestDto
	err := c.Bind(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResError(err.Error()))
		return
	}

	companies, aErr := ct.service.RegisterCompany(dto)
	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(companies)
	c.JSON(http.StatusOK, response)
}
