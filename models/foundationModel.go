package models

import (
	"time"

	"gorm.io/gorm"
)

type Foundation struct {
	UserID     uint           `gorm:"unique;primaryKey"`
	LicenseID uint            `gorm:"unique" json:"license_id"`
	AddressID  uint           `json:"address_id"`
	Address    Address        `gorm:"foreignKey:AddressID"`
	User       User           `gorm:"foreignKey:UserID"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type RegisterFoundationAPI struct {
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	LincenseID  uint           `json:"license_id"`
	AddressName string         `json:"address_name"`
	Latitude    string         `json:"lat"`
	Longitude   string         `json:"long"`
	CityID      uint           `json:"city_id"`
	ProvinceID  uint           `json:"province_id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type FoundationAPI struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
