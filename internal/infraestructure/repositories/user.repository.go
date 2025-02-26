package repositories

import (
	"col-moda/internal/domain/entities"

	"gorm.io/gorm"
)

type UserRepository struct {
	BaseRepository
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		BaseRepository: BaseRepository{
			db: db,
		},
	}
}

func (r UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	e := r.db.Where("email = ?", email).First(&user)

	if e.Error != nil {
		if e.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, e.Error
	}

	return &user, nil
}

func (r UserRepository) CreateUser(user entities.User) (*bool, error) {
	result := r.db.Create(&user)

	if result.Error != nil {
		return nil, result.Error
	}

	success := result.RowsAffected > 0

	return &success, nil
}
