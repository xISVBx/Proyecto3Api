package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Company struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primaryKey"`
	CompanyName string
	Description string
}

