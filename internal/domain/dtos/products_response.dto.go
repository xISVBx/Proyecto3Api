package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type ProductResponseDto struct {
	ID               uuid.UUID  `json:"id"`
	Name             string     `json:"name"`
	ShortDescription string     `json:"short_description"`
	LongDescription  string     `json:"long_description"`
	Price            float64    `json:"price"`
	Stock            int        `json:"stock"`
	CompanyID        uuid.UUID  `json:"company_id"`
	CategoryID       uuid.UUID  `json:"category_id"`
	SubCategoryID    uuid.UUID  `json:"sub_category_id"`
	DiscountID       *uuid.UUID `json:"discount_id,omitempty"`
	TagID            uuid.UUID  `json:"tag_id"`
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
