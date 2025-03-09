package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type CityService struct {
	uow *repositories.UnitOfWork
}

func NewCityService(uow *repositories.UnitOfWork) *CityService {
	return &CityService{
		uow: uow,
	}
}

func (s CityService) FindCitiesByFilters(cityRequest dtos.CityRequestDto) ([]dtos.CityResponseDto, *models.AppError) {
	
	s.uow.Begin()

	defer s.uow.Commit()

	cities, err := s.uow.CityRepo.FindCitiesByFilters(cityRequest)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	citiesResponse := dtos.CitiesResponseFromEntities(cities)

	return citiesResponse, nil
}
