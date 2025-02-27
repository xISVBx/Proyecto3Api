package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductView struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID int
	UserID    int

	Product Product `gorm:"foreignKey:ProductID"`
	User    User `gorm:"foreignKey:UserID"`
}
