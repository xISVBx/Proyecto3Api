package repositories

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/entities"

	"gorm.io/gorm"
)

type DepartmentRepository struct {
    db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
    return &DepartmentRepository{db: db}
}

func (r DepartmentRepository) FindAllDepartmentsByFilters(dto dtos.DepartmentRequest) ([]entities.Department, error) {
    var departments []entities.Department
    var query = r.db
    if(dto.DepartmentID != nil) {
        query = query.Where("id = ?", *dto.DepartmentID)
    }
    if dto.DepartmentName != "" {
        query = query.Where("LOWER(name) LIKE LOWER(?)", "%"+dto.DepartmentName+"%")
    }
    err := query.Find(&departments).Error
    if err != nil {
        return nil, err
    }

    return departments, nil
}