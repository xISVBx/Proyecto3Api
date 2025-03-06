package repositories

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
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

// FindCompaniesByFilters busca empresas aplicando filtros opcionales
func (r *CompanyRepository) FindCompaniesByFilters(dto dtos.CompanyRequestDto) ([]entities.Company, error) {
	var companies []entities.Company
	query := r.db.Model(&entities.Company{})

	// Aplica filtros si existen
	if dto.ID != uuid.Nil {
		query = query.Where("id = ?", dto.ID)
	}
	if dto.CompanyName != "" {
		query = query.Where("company_name ILIKE ?", "%"+dto.CompanyName+"%") // Búsqueda flexible
	}

	// Ejecuta la consulta
	if err := query.Find(&companies).Error; err != nil {
		return nil, err
	}

	return companies, nil
}
