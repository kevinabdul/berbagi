package models

import (
	"time"
)

type ConfirmService struct {
	VolunteerID uint      `gorm:"primaryKey"`
	AddressID   uint      `gorm:"not null" json:"recipient_address_id" form:"recipient_address_id"`
	StartDate   time.Time `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate  time.Time `gorm:"not null" json:"finish_date" form:"finish_date"`
}

type DataProduct struct {
	InvoiceID  uint `json:"invoice_id"`
	DonorID    uint `json:"donor_id"`
	ProductID  uint `json:"product_id" form:"product_id"`
	AddressID  uint `json:"recipient_address_id" form:"recipient_address_id"`
	Quantity   int  `json:"quantity" form:"quantity"`
	TotalPrice int  `json:"total_price" `
}

type ResponseConfirmService struct {
	Message string
	Data    DataService
}

type DataService struct {
	InvoiceID   uint      `json:"invoice_id"`
	VolunteerID uint      `json:"volunteer_id"`
	AddressID   uint      `json:"recipient_address_id" form:"recipient_address_id"`
	StartDate   time.Time `json:"start_date" form:"start_date"`
	FinishDate  time.Time `json:"finish_date" form:"finish_date"`
}
