package entities

import (
	"time"

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
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Order   Order   `gorm:"foreignKey:OrderID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

