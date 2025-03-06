package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type CategoryResponseDto struct {
	ID          uuid.UUID
	Name        string
	Descripcion string
}

func CategoriesResponseFromEntitie(categories []entities.Category) []CategoryResponseDto {
	var categoriesResponse []CategoryResponseDto
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, CategoryResponseDto{
			ID:          category.ID,
			Name:        category.Name,
			Descripcion: category.Description,
		})
	}
	return categoriesResponse
}
