package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type ProductService struct {
	uow *repositories.UnitOfWork
}

func NewProductService(uow *repositories.UnitOfWork) *ProductService {
	return &ProductService{
		uow: uow,
	}
}

func (s ProductService) FindProductsByFilters(productsRequest dtos.ProductRequestDto) ([]dtos.ProductResponseDto, *models.AppError) {
	
	s.uow.Begin()

	defer s.uow.Commit()
	
	products, err := s.uow.ProductRepo.FindProductsByFilters(productsRequest)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	productsResponse := dtos.ProductsResponseDtoFromEntities(products)

	return productsResponse, nil
}
