package repositories

import (
	"col-moda/internal/domain/entities"

	"gorm.io/gorm"
)

type ProductRepository struct {
	BaseRepository
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
}

func (r RoleRepository) FindProductsByFilters(searchQuery string) ([]entities.Product, error) {
	var products []entities.Product

	query := `
        SELECT * FROM products 
        WHERE to_tsvector(name || ' ' || description) @@ plainto_tsquery(?)
        ORDER BY ts_rank(to_tsvector(name || ' ' || description), plainto_tsquery(?)) DESC`

	e := r.db.Raw(query, searchQuery, searchQuery).Scan(&products)

	if e.Error != nil {
		return nil, e.Error
	}

	return products, nil
}
