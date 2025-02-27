package dtos

import (
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
)

type CategoryResponse struct {
	ID          uuid.UUID
	Name        string
	Descripcion string
}

func CategoriesResponseFromCategories(categories []entities.Category) []CategoryResponse {
	var categoriesResponse []CategoryResponse
	for _, category := range categories {
		categoriesResponse = append(categoriesResponse, CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Descripcion: category.Description,
		})
	}
	return categoriesResponse
}
