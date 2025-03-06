package controllers

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productS *services.ProductService
}

func NewProductController(s *services.ProductService) *ProductController {
	return &ProductController{
		productS: s,
	}
}

// FindProductsByFilters Obtener los productos por filtros.
// @Summary Obtener todos los productos por filtros
// @Description Retorna un listado de los productos aplicando filtros.
// @Tags Products
// @Accept json
// @Produce json
// @Param SearchQuery query string false "Texto de búsqueda"
// @Success 200 {object} models.AppResponse{data=[]dtos.ProductResponseDto}
// @Failure 400 {object} models.AppResponse{data=interface{}}
// @Failure 409 {object} models.AppResponse{data=interface{}}
// @Router /api/v1/products [get]
func (ct ProductController) FindProductsByFilters(c *gin.Context) {

	var dto dtos.ProductRequestDto

	err := c.Bind(&dto)

	if err != nil {
		c.JSON(http.StatusBadRequest, models.ResError("Invalid data"))
		return
	}

	products, aErr := ct.productS.FindProductsByFilters(dto)

	if aErr != nil {
		c.JSON(aErr.Status, models.ResFromApp(aErr))
		return
	}

	response := models.ResSuccess(products)
	c.JSON(http.StatusOK, response)
}
