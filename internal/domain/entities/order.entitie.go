package entities

import (
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

	User User `gorm:"foreignKey:UserID"`
}
