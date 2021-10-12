package models

import (
	"time"

	"gorm.io/gorm"
)

type Proficiency struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"unique;type:varchar(50);not null" json:"name" form:"name"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
