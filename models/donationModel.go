package models

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	ID            uint           `gorm:"primaryKey" json:"donation_id" form:"donation_id"`
	DonorID       uint           `json:"donor_id" form:"donor_id"`
	RecipientID   uint           `json:"recipient_id" form:"recipient_id"` // as prevention if donation can be given to Children
	RequestID     uint           `json:"request_id" form:"request_id"`
	Amount        int            `json:"amount" form:"amount"`
	PaymentStatus string         `sql:"type:ENUM('true', 'false')" gorm:"default:false" json:"payment_status" form:"payment_status"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type DonationResponse struct {
	DonationID    uint      `json:"donation_id" form:"donation_id"`
	DonorID       uint      `json:"donor_id" form:"donor_id"`
	RecipientID   uint      `json:"recipient_id" form:"recipient_id"`
	RequestID     uint      `json:"request_id" form:"request_id"`
	Amount        int       `json:"amount" form:"amount"`
	PaymentStatus string    `json:"payment_status" form:"payment_status"`
	MadeAt        time.Time `json:"made_at"`
}

type NewDonationResponseAPI struct {
	DonationID  uint `json:"donation_id"`
	DonorID     uint `json:"donor_id"`
	RecipientID uint `json:"recipient_id"`
	RequestID   uint `json:"request_id"`
	Amount      int  `json:"amount"`
}

type DonationCart struct {
	DonorID     uint           `gorm:"uniqueIndex:composite" json:"donor_id"`
	RecipientID uint           `gorm:"uniqueIndex:composite" json:"recipient_id" form:"recipient_id"`
	RequestID   uint           `gorm:"uniqueIndex:composite" json:"request_id" form:"request_id"`
	Amount      int            `json:"amount" form:"amount"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	User        User           `gorm:"foreignKey:RecipientID" json:"-"`
	// Request     Request        `gorm:"foreignKey:RequestID" json:"-"`
}

// For read, update, delete operation in donation_cart
type DonationInputData struct {
	DonorID     uint `json:"donor_id" form:"donor_id"`
	RecipientID uint `json:"recipient_id" form:"recipient_id"`
	RequestID   uint `json:"request_id" form:"request_id"`
	Amount      int  `json:"amount" form:"amount"`
	PaymentID   int  `json:"payment_id" form:"payment_id"`
}
