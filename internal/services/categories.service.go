package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type CategorieService struct {
	roleR *repositories.CategoryRepository
}

func NewCategorieService(r *repositories.CategoryRepository) *CategorieService {
	return &CategorieService{
		roleR: r,
	}
}

func (s CategorieService) FindAllCategories() ([]dtos.CategoryResponseDto, *models.AppError) {
	categories, err := s.roleR.FindAllCategories()

	if err != nil {
		return nil, models.NewServerError(err)
	}

	categoryDto := dtos.CategoriesResponseFromEntitie(categories)

	return categoryDto, nil
}
