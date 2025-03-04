package repositories

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/entities"

	"gorm.io/gorm"
)

type CityRepository struct {
	BaseRepository
}

func NewCityRepository(db *gorm.DB) *CityRepository {
	return &CityRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
}

func (r CityRepository) FindCitiesByFilters(cityRequest dtos.CityRequest) ([]entities.City, error) {
	var cities []entities.City
	query := r.db

	if cityRequest.CityID != nil {
		query = query.Where("id = ?", *cityRequest.CityID)
	}

	if cityRequest.DepartmentID != nil {
		query = query.Where("department_id = ?", cityRequest.DepartmentID)
	}

	if cityRequest.CityName != "" {
		query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+cityRequest.CityName+"%")
	}

	err := query.Find(&cities).Error
	if err != nil {
		return nil, err
	}

	return cities, nil
}
