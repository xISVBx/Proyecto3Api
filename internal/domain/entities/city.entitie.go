package entities

import (
)

type City struct {
	ID           int    `gorm:"primaryKey"`
	Name         string `gorm:"not null"`
	DepartmentID int    `gorm:"not null"`

	Department Department `gorm:"foreignKey:DepartmentID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}