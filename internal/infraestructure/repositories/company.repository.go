package repositories

import (
	"col-moda/internal/domain/entities"

	"gorm.io/gorm"
)

type CompanyRepository struct {
	BaseRepository
}

func NewCompanyRepository(db *gorm.DB) *CompanyRepository {
	return &CompanyRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
}

func (r CategoryRepository) FindAllCompanies() ([]entities.Category, error) {
	var categories []entities.Category
	e := r.db.Find(&categories)

	if e.Error != nil {
		return nil, e.Error
	}

	return categories, nil
}
