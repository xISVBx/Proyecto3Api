package repositories

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/entities"
	"fmt"

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

func (r ProductRepository) FindProductsByFilters(productsRequest dtos.ProductRequest) ([]entities.Product, error) {
	var products []entities.Product

	query := `
    SELECT * FROM products 
    WHERE to_tsvector(name || ' ' || short_description || ' ' || long_description) @@ to_tsquery(? || ':*')
    ORDER BY ts_rank(to_tsvector(name || ' ' || short_description || ' ' || long_description), to_tsquery(? || ':*')) DESC`

	e := r.db.Raw(query, productsRequest.SearchQuery, productsRequest.SearchQuery).Scan(&products)

	if e.Error != nil {
		return nil, e.Error
	}

	fmt.Println("Parámetro de búsqueda")
	fmt.Println(productsRequest.SearchQuery)

	if len(products) != 0 {
		fmt.Println("Resultado")
		fmt.Println(products[0].Name)
	}

	return products, nil
}
