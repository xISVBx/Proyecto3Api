package repositories

import (
	"col-moda/internal/domain/entities"

	"gorm.io/gorm"
)

type CategoryRepository struct {
	BaseRepository
}

func NewCategoriesRepository(db *gorm.DB) *CategoryRepository {
	return &CategoryRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
}

func (r CategoryRepository) FindAllCategories() ([]entities.Category, error) {
	var categories []entities.Category
	e := r.db.Find(&categories)

	if e.Error != nil {
		return nil, e.Error
	}

	return categories, nil
}
