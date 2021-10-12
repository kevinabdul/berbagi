package models

import (
	"time"

	"gorm.io/gorm"
)

type Donor struct {
	UserID    uint           `gorm:"unique;primaryKey"`
	BirthDate string         `json:"birth_date"`
	AddressID uint           `json:"address_id"`
	Address   Address        `gorm:"foreignKey:AddressID"`
	User      User           `gorm:"foreignKey:UserID"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type RegisterDonorAPI struct {
	Name        string         `json:"name"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	NIK         string         `json:"nik"`
	BirthDate   string         `json:"birth_date"`
	AddressName string         `json:"address_name"`
	Latitude    string         `json:"lat"`
	Longitude   string         `json:"long"`
	CityID      uint           `json:"city_id"`
	ProvinceID  uint           `json:"province_id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type DonorAPI struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
