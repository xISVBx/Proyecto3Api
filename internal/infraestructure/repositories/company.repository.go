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

// CreateCompany guarda una nueva empresa en la base de datos
func (r *CompanyRepository) CreateCompany(company entities.Company) (*bool, error) {
	result := r.db.Create(&company)

	if result.Error != nil {
		return nil, result.Error
	}

	success := result.RowsAffected > 0

	return &success, nil
}

// CreateCompany guarda una nueva empresa en la base de datos
func (r *CompanyRepository) CreateUserCompany(uCompany *entities.UserCompany) (*bool, error) {
	result := r.db.Create(&uCompany)

	if result.Error != nil {
		return nil, result.Error
	}

	success := result.RowsAffected > 0

	return &success, nil
}
