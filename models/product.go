package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"not null" json:"name" form:"name"`
	Price       int            `gorm:"not null" json:"price" form:"price"`
	Description string         `json:"description" form:"description"`
	CategoryID  uint           `gorm:"not null" json:"category_id" form:"category_id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
}
