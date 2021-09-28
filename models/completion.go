package models

import (
	"time"

	"gorm.io/gorm"
)

type Completion struct {
	VolunteerID uint           `gorm:"primaryKey"`
	InvoiceID   uint           `json:"invoice_id" form:"invoice_id"`
	Status      string         `json:"status" form:"status"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Volunteer   Volunteer      `gorm:"foreignKey:VolunteerID"`
}

type CompletionDetail struct {
	InvoiceID  uint           `json:"invoice_id" form:"invoice_id"`
	UserID     uint           `json:"recipient_id" form:"recipient_id"`
	AddressID  uint           `gorm:"not null" json:"recipient_address_id" form:"recipient_address_id"`
	StartDate  time.Time      `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate time.Time      `gorm:"not null" json:"finish_date" form:"finish_date"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
	Address    Address        `gorm:"foreignKey:AddressID"`
	User       User           `gorm:"foreignKey:UserID"`
}
