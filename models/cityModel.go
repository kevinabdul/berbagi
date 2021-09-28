package models

import (
	"time"

	"gorm.io/gorm"
)

type City struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(50)" json:"name"`
	ProvinceID uint           `json:"province_id"`
	Province   Province       `gorm:"foreignKey:ProvinceID"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
