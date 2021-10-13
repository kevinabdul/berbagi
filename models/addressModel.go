package models

import (
	"time"

	"gorm.io/gorm"
)

type Address struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	Name       string         `gorm:"type:varchar(150)" json:"name"`
	Latitude   string         `gorm:"type:varchar(150)" json:"latitude"`
	Longitude  string         `gorm:"type:varchar(150)" json:"longitude"`
	CityID     uint           `json:"city_id"`
	ProvinceID uint           `json:"province_id"`
	City       City           `gorm:"foreignKey:CityID"`
	Province   Province       `gorm:"foreignKey:ProvinceID"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type NearestAddressIdResponseAPI struct {
	ID       uint    `json:"address_id"`
	Distance float64 `json:"distance"`
}

type LocationPointResponseAPI struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type NearbyInputData struct {
	UserID      uint
	Role        string
	Latitude    float64
	Longitude   float64
	Range       float64
	GetResource string
	Type        string
}
