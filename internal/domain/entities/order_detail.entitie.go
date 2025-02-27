package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type OrderDetail struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID       uuid.UUID
	ProductID     uuid.UUID
	Quantity      int
	PricePerUnit  float64
	PriceFinal    float64
	TotalDiscount float64

	Order   Order   `gorm:"foreignKey:OrderID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

