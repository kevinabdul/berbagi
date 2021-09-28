package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	DonorID         uint           `gorm:"primaryKey"`
	InvoiceID       uint           `json:"invoice_id" form:"invoice_id"`
	PaymentMethodID uint           `json:"payment_id" form:"payment_id"`
	StatusPayment   string         `json:"payment_status" form:"payment_status"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index"`
	Donor           Donor          `gorm:"foreignKey:DonorID"`
	Product         Product        `gorm:"foreignKey:ProductID"`
}

type TransactionDetail struct {
	InvoiceID uint           `json:"invoice_id" form:"invoice_id"`
	UserID    uint           `json:"recipient_id" form:"recipient_id"`
	ProductID uint           `gorm:"not null;primaryKey" json:"product_id" form:"product_id"`
	Quantity  int            `gorm:"not null" json:"quantity" form:"quantity"`
	BuyPrice  int            `gorm:"not null" json:"buy_price" form:"buy_price"`
	AddressID uint           `json:"recipient_address_id" form:"recipient_address_id"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Product   Product        `gorm:"foreignKey:ProductID"`
	Address   Address        `gorm:"foreignKey:AddressID"`
	User      User           `gorm:"foreignKey:UserID"`
}
