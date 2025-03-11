package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type CategoryResponseDto struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}


func CategoriesResponseFromEntitie(categories []entities.Category) []CategoryResponseDto {
	var categoriesResponse []CategoryResponseDto
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, CategoryResponseDto{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
	return categoriesResponse
}
