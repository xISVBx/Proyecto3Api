package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID               uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name             string
	ShortDescription string
	LongDescription  string
	Price            float64
	Stock            int
	CompanyID        uuid.UUID
	CategoryID       uuid.UUID
	SubCategoryID    uuid.UUID
	DiscountID       *uuid.UUID `gorm:"type:uuid"` // ✅ Permite NULL
	TagID            uuid.UUID
	CreatedAt        time.Time
	UpdatedAt        time.Time

	Company        Company          `gorm:"foreignKey:CompanyID"`
	Category       Category         `gorm:"foreignKey:CategoryID"`
	SubCategory    SubCategory      `gorm:"foreignKey:SubCategoryID"`
	Tag            Tag              `gorm:"foreignKey:TagID"`
	ImagesProducts []ImageProduct `gorm:"foreignKey:ProductID"`
	Discount       *Discount        `gorm:"foreignKey:DiscountID"` // ✅ Relación opcional
}
