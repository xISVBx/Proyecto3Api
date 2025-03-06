package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type ProductResponseDto struct {
	ID               uuid.UUID
	Name             string
	ShortDescription string
	LongDescription  string
	Price            float64
	Stock            int
	CompanyID        uuid.UUID
	CategoryID       uuid.UUID
	SubCategoryID    uuid.UUID
	DiscountID       *uuid.UUID
	TagID            uuid.UUID
}

func ProductsResponseDtoFromEntities(products []entities.Product) []ProductResponseDto {
	var productsResponse []ProductResponseDto
	for _, product := range products {
		productsResponse = append(productsResponse, ProductResponseDto{
			ID:               product.ID,
			Name:             product.Name,
			ShortDescription: product.ShortDescription,
			LongDescription:  product.LongDescription,
			Price:            product.Price,
			Stock:            product.Stock,
			CompanyID:        product.CompanyID,
			SubCategoryID:    product.SubCategoryID,
			DiscountID:       product.DiscountID,
			TagID:            product.TagID,
		})
	}
	return productsResponse
}
