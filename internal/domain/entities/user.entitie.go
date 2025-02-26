package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name      string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
	CityID    int
	Address   string
	Phone     string
	RoleID    uuid.UUID

	City City `gorm:"foreignKey:CityID"`
	Role Role `gorm:"foreignKey:RoleID"`
}
