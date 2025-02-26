package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Discount struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:uuid;primaryKey"`
	Type         string
	DiscountValue float64
	StartDate     time.Time
	EndDate       time.Time
	MinQuantity   int
	MaxQuantity   int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

