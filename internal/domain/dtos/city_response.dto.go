package dtos

import "col-moda/internal/domain/entities"

type CityResponse struct {
	ID          int    
	DepartmentID int    
	CityName    string 
}

func CitiesResponseFromEntities(e []entities.City) []CityResponse {
	var cityResponse []CityResponse
	for _, city := range e {
		cityResponse = append(cityResponse, CityResponse{
			ID:          city.ID,
			DepartmentID: city.DepartmentID,
			CityName:    city.Name,
		})
	}
	return cityResponse;
}	