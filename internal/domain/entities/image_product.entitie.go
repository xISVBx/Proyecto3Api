package entities

import "github.com/google/uuid"

type ImageProduct struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey"`
	ProductID int
	ImagePath string

	Product Product `gorm:"foreignKey:ProductID"`
}
