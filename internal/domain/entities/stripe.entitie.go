package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Stripe struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primaryKey"`
	OrderID       uuid.UUID
	Amount        float64
	State         string
	PaymentMethod string
	Currency      string
	CreatedAt     time.Time
	UpdatedAt     time.Time

	Order Order `gorm:"foreignKey:OrderID"`
}