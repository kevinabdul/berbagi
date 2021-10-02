package models

import (
	"time"

	"gorm.io/gorm"
)

type ServiceCart struct {
	VolunteerID uint           `json:"volunteer_id"`
	UserID      uint           `json:"recipient_id" form:"recipient_id"`
	AddressID   uint           `gorm:"not null" json:"recipient_address_id" form:"recipient_address_id"`
	StartDate   time.Time      `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate  time.Time      `gorm:"not null" json:"finish_date" form:"finish_date"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Address     Address        `gorm:"foreignKey:AddressID" json:"-"`
	User        User           `gorm:"foreignKey:UserID" json:"-"`
}

type InputService struct {
	UserID     uint   `gorm:"not null" json:"recipient_id" form:"recipient_id"`
	StartDate  string `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate string `gorm:"not null" json:"finish_date" form:"finish_date"`
}

type ResponseService struct {
	VolunteerName string    `json:"volunteer_name"`
	UserName      string    `json:"recipient_name" `
	AddressName   string    `json:"recipient_address" `
	StartDate     time.Time `json:"start_date" `
	FinishDate    time.Time `json:"finish_date" `
}
