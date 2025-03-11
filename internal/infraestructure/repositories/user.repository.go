package repositories

import (
	"col-moda/internal/domain/dtos"
	"col-moda/internal/domain/entities"

	"github.com/google/uuid"
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

func (r UserRepository) FindUserWithCompany(userID uuid.UUID) (*dtos.UserCompanyResponseDto, error) {
	var userCompany dtos.UserCompanyResponseDto

	e := r.db.Raw(`
		SELECT 
			u.id, u.name, u.last_name, u.email, u.phone,
			uc.company_id, c.company_name, u.role_id
		FROM public.users u
		INNER JOIN public.user_companies uc ON u.id = uc.user_id
		INNER JOIN public.companies c ON uc.company_id = c.id
		WHERE u.id = ?
	`, userID).Scan(&userCompany)

	if e.Error != nil {
		return nil, e.Error
	}

	return &userCompany, nil
}
