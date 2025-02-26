package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubCategory struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CategoryID  uuid.UUID
	Description string

	Category Category `gorm:"foreignKey:CategoryID"`
}