package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type ProductResponse struct {
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

func ProductsResponseFromEntities(products []entities.Product) []ProductResponse {
	var productsResponse []ProductResponse
	for _, product := range products {
		productsResponse = append(productsResponse, ProductResponse{
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
