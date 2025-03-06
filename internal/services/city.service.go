package services

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/models"
	"col-moda/internal/infraestructure/repositories"
)

type CityService struct {
	cityR *repositories.CityRepository
}

func NewCityService(r *repositories.CityRepository) *CityService {
	return &CityService{
		cityR: r,
	}
}

func (s CityService) FindCitiesByFilters(cityRequest dtos.CityRequestDto) ([]dtos.CityResponseDto, *models.AppError) {
	cities, err := s.cityR.FindCitiesByFilters(cityRequest)

	if err != nil {
		return nil, models.NewServerError(err)
	}

	citiesResponse := dtos.CitiesResponseFromEntities(cities)

	return citiesResponse, nil
}
