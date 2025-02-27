package entities

import "github.com/google/uuid"

type ShoppingCart struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	UserID    uuid.UUID
	ProductID uuid.UUID
	Quantity  int

	Product Product `gorm:"foreignKey:ProductID"`
	User User `gorm:"foreignKey:UserID"`
}
