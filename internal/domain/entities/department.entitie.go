package entities

import (
)

type Department struct {
	ID   int    `gorm:"primaryKey"`
	Name string `gorm:"not null"`

	Cities []City `gorm:"foreignKey:DepartmentID"`
}
