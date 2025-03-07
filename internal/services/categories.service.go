package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type CategorieService struct {
	uow *repositories.UnitOfWork
}

func NewCategorieService(uow *repositories.UnitOfWork) *CategorieService {
	return &CategorieService{
		uow: uow,
	}
}

func (s CategorieService) FindAllCategories() ([]dtos.CategoryResponseDto, *models.AppError) {
	categories, err := s.uow.CategoryRepo.FindAllCategories()

	if err != nil {
		return nil, models.NewServerError(err)
	}

	categoryDto := dtos.CategoriesResponseFromEntitie(categories)

	return categoryDto, nil
}
