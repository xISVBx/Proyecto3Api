package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type ProductService struct {
	productR *repositories.ProductRepository
}

func NewProductService(r *repositories.ProductRepository) *ProductService {
	return &ProductService{
		productR: r,
	}
}

func (s ProductService) FindProductsByFilters(productsRequest dtos.ProductRequest) ([]dtos.ProductResponse, *models.AppError) {
	products, err := s.productR.FindProductsByFilters(productsRequest)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	productsResponse := dtos.ProductsResponseFromEntities(products)

	return productsResponse, nil
}
