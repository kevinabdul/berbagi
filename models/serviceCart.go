package models

import (
	"time"

	"gorm.io/gorm"
)

type ServiceCart struct {
	VolunteerID uint           `gorm:"primaryKey"`
	UserID      uint           `json:"recipient_id" form:"recipient_id"`
	AddressID   uint           `gorm:"not null" json:"recipient_address_id" form:"recipient_address_id"`
	StartDate   time.Time      `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate  time.Time      `gorm:"not null" json:"finish_date" form:"finish_date"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Volunteer   Volunteer      `gorm:"foreignKey:VolunteerID"`
	Address     Address        `gorm:"foreignKey:AddressID"`
	User        User           `gorm:"foreignKey:UserID"`
}