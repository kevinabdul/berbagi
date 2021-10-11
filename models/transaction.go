package models

import (
	"time"

	"gorm.io/gorm"
)

type Transaction struct {
	DonorID         uint           `gorm:"primaryKey" json:"donor_id"`
	InvoiceID       string         `gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	PaymentMethodID uint           `json:"payment_id" form:"payment_id"`
	PaymentStatus   string         `gorm:"type:enum('pending', 'expired', 'cancelled', 'paid');default:'pending'" json:"payment_status"`
	Total           int            `gorm:"type:int" json:"total"`
	CreatedAt       time.Time      `json:"-"`
	UpdatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
	Donor           Donor          `gorm:"foreignKey:DonorID"`
	PaymentMethod   PaymentMethod  `gorm:"foreignKey:PaymentMethodID"`
}

type TransactionDetail struct {
	InvoiceID        string `gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	RecipientID      uint   `gorm:"primaryKey" json:"recipient_id"`
	ProductPackageID uint   `gorm:"primaryKey;type:uint" json:"product_package_id"`
	PackagePrice     int    `json:"package_price"`
	Quantity         uint   `gorm:"not null;type:smallint" json:"quantity"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Transaction      Transaction    `gorm:"foreignKey:InvoiceID;references:InvoiceID"`
	ProductPackage   ProductPackage `gorm:"foreignKey:ProductPackageID"`
	Children         Children       `gorm:"foreignKey:RecipientID"`
}

type TransactionDonationDetail struct {
	InvoiceID   string `gorm:"primaryKey;not null;type:varchar(60)" json:"invoice_id"`
	DonationID  uint   `gorm:"primaryKey;type:uint" json:"donation_id"`
	RecipientID uint   `gorm:"primaryKey" json:"recipient_id"`
	Amount      int    `json:"amount"`
	Purpose     string `json:"purpose"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Transaction Transaction `gorm:"foreignKey:InvoiceID;references:InvoiceID"`
	Donation    Donation    `gorm:"foreignKey:DonationID"`
	Foundation  Foundation  `gorm:"foreignKey:RecipientID"`
}

// Response struct used in case of a succesful checkout in post checkout endpoint
// Succesful checkout means we are able to delete data from carts table, creating new data in transactions table,
// and moving the deleted data into transaction_details table. Any failure in those step will fail whole transaction
type TransactionAPI struct {
	InvoiceID       string `json:"invoice_id"`
	Total           int    `json:"total"`
	PaymentMethodID uint   `json:"payment_method_id"`
	Description     string `json:"description"`
}

// Commonly used when user try to do a payment.
// AddPaymentByUserId will try to find corresponding transaction in a database based on UserId and information provided in UserPaymentAPI struct.
// This struct will be used as a placeholder of above query result.
type TransactionDetailAPI struct {
	DonorID           string    `json:"-"`
	InvoiceID         string    `json:"-"`
	RecipientID       uint      `json:"recipient_id"`
	PaymentStatus     string    `json:"payment_status"`
	ProductPackageID  uint      `json:"product_package_id"`
	PackagePrice      int       `json:"package_price"`
	Quantity          uint      `json:"quantity"`
	Total             int       `json:"total"`
	PaymentMethodID   uint      `json:"-"`
	PaymentMethodName string    `json:"-"`
	Description       string    `json:"-"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}

type TransactionDonationDetailAPI struct {
	DonorID           string    `json:"-"`
	DonationID        uint      `json:"-"`
	InvoiceID         string    `json:"-"`
	RecipientID       uint      `json:"recipient_id"`
	PaymentStatus     string    `json:"payment_status"`
	Purpose           string    `json:"purpose"`
	Amount            int       `json:"amount"`
	PaymentMethodID   uint      `json:"-"`
	PaymentMethodName string    `json:"-"`
	Description       string    `json:"-"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
}
