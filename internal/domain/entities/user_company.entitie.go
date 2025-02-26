package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserCompany struct {
	gorm.Model
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey"`
	CompanyID uuid.UUID `gorm:"type:uuid;primaryKey"`

	User    User    `gorm:"foreignKey:UserID"`
	Company Company `gorm:"foreignKey:CompanyID"`
}

