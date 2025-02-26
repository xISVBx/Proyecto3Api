package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ID              uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID          uuid.UUID
	TotalAmount     float64
	Status         string
	TotalDiscount  float64
	TotalFinal     float64
	ShippingCost   float64
	ShippingAddress string
	CreatedAt      time.Time
	UpdatedAt      time.Time

	User User `gorm:"foreignKey:UserID"`
}
