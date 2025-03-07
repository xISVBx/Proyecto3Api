package repositories

import (
	"col-moda/internal/domain/entities"

	"gorm.io/gorm"
)

type RoleRepository struct {
	BaseRepository
}

func NewRoleRepository(db *gorm.DB) *RoleRepository {
	return &RoleRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
}

func (r RoleRepository) FindAllRoles() ([]entities.Role, error) {
	var roles []entities.Role
	e := r.db.Find(&roles)

	if e.Error != nil {
		return nil, e.Error
	}

	return roles, nil
}

func (r RoleRepository) FindUserRole() (*entities.Role, error) {
	var role entities.Role
	e := r.db.Where("description = ?", "Usuario").First(&role)

	if e.Error != nil {
		if e.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, e.Error
	}

	return &role, nil
}

func (r RoleRepository) FindCompanyRole() (*entities.Role, error) {
	var role entities.Role
	e := r.db.Where("description = ?", "Company").First(&role)

	if e.Error != nil {
		if e.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, e.Error
	}

	return &role, nil
}