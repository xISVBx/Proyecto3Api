package dtos

import "col-moda/internal/domain/entities"

type CityResponseDto struct {
	ID           int    `json:"id"`
	DepartmentID int    `json:"department_id"`
	CityName     string `json:"city_name"`
}

func CitiesResponseFromEntities(e []entities.City) []CityResponseDto {
	var cityResponse []CityResponseDto
	for _, city := range e {
		cityResponse = append(cityResponse, CityResponseDto{
			ID:           city.ID,
			DepartmentID: city.DepartmentID,
			CityName:     city.Name,
		})
	}
	return cityResponse
}
