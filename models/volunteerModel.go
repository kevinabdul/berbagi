package models

import (
	"time"

	"gorm.io/gorm"
)

type Volunteer struct {
	UserID        uint           `gorm:"not null" json:"user_id" form:"user_id"`
	BirthDate     string         `json:"birth_date" form:"birth_date"`
	ProficiencyID uint           `gorm:"not null" json:"proficiency_id" form:"proficiency_id"`
	AddressID     uint           `gorm:"not null" json:"address_id" form:"address_id"`
	Address       Address        `gorm:"foreignKey:AddressID"`
	User          User           `gorm:"foreignKey:UserID"`
	Proficiency   Proficiency    `gorm:"foreignKey:ProficiencyID"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}

type ProfileVolunteerAPI struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	NIK             string `json:"nik"`
	BirthDate       string `json:"birth_date"`
	AddressName     string `json:"address_name"`
	CityName        string `json:"city_name"`
	ProvinceName    string `json:"province_name"`
	ProficiencyName string `json:"proficiency_name"`
}

type VolunteerAPI struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
