package models

import (
	"time"

	"gorm.io/gorm"
)

type Volunteer struct {
	ID            uint           `gorm:"primaryKey"`
	UserID        uint           `gorm:"not null" json:"user_id" form:"user_id"`
	BirthDate     string         `json:"birth_date" form:"birth_date"`
	ProficiencyID uint           `gorm:"not null" json:"proficiency_id" form:"proficiency_id"`
	AddressID     uint           `gorm:"not null" json:"address_id" form:"address_id"`
	Address       Address        `gorm:"foreignKey:AddressID"`
	User          User           `gorm:"foreignKey:UserID"`
	Proficiency   Proficiency    `gorm:"foreignKey:ProficiencyID"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type RegisterVolunteerAPI struct {
	Name          string `json:"name"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	NIK           string `json:"nik"`
	BirthDate     string `json:"birth_date"`
	AddressName   string `json:"address_name"`
	Latitude      string `json:"lat"`
	Longitude     string `json:"long"`
	CityID        uint   `json:"city_id"`
	ProvinceID    uint   `json:"province_id"`
	ProficiencyID uint   `json:"proficiency_id"`
}

type VolunteerAPI struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
