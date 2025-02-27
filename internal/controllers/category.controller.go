package controllers

import (
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryS *services.CategorieService
}

func NewCategoryController(s *services.CategorieService) *CategoryController {
	return &CategoryController{
		categoryS: s,
	}
}

// FindAllCategories Obtener todas las categorias de la aplicacion
// @Summary Obtener todas las categorias
// @Description Retorna un listado de las categorias de la aplicacion.
// @Tags Categories
// @Accept json
// @Produce json
// @Success 201 {object} models.AppResponse{data=[]dtos.CategoryResponse}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 409 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/categories [get]
func (ct CategoryController) FindAllCategories(c *gin.Context) {

	categories, aErr := ct.categoryS.FindAllCategories()

	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(categories)
	c.JSON(http.StatusOK, response)
}
