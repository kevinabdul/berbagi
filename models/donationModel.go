package models

import (
	"time"

	"gorm.io/gorm"
)

type Donation struct {
	DonationID    uint
	DonorID       uint
	RecipientID   uint // as prevention if donation can be given to Children
	RequestID     uint
	Amount        float64
	PaymentStatus bool
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

type DonationResponse struct {
	DonationID  uint
	DonorID     uint
	RecipientID uint
	RequestID   uint
	Amount      float64
	PaymentStatus bool
	MadeAt      time.Time
}

type NewDonation struct {
	DonorID     uint
	RecipientID uint
	RequestID   uint
	Amount      float64
}

type NewDonationResponseAPI struct {
	DonationID  uint `json:"donation_id"`
	DonorID     uint `json:"donor_id"`
	RecipientID uint `json:"recipient_id"`
	// AddressID   uint    `json:"address_id"`
	RequestID uint    `json:"request_id"`
	Amount    float64 `json:"amount"`
}

type DonationCart struct {
	DonorID     uint           `json:"donor_id"`
	RecipientID uint           `json:"recipient_id" form:"recipient_id"`
	// AddressID   uint           `gorm:"not null" json:"recipient_address_id" form:"recipient_address_id"`
	RequestID   uint           `json:"request_id" form:"request_id"`
	Amount      float64        `json:"amount" form:"amount"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Address     Address        `gorm:"foreignKey:AddressID" json:"-"`
	User        User           `gorm:"foreignKey:RecipientID" json:"-"`
	Request     Request        `gorm:"foreignKey:RequestID" json:"-"`
}

// For read, update, delete operation in donation_cart
type CartItemInputData struct {
	DonorID     uint
	RecipientID uint
	RequestID   uint
	Amount      float64
}

type DonationCheckout struct {
	DonorID     uint
	RecipientID uint
	RequestID   uint
}

// type UnresolvedDonation struct {
// 	DonationID          uint
// 	DonorID     uint
// 	RecipientID uint
// 	AddressID   uint
// 	RequestID   uint
// 	Amount      float64
// 	CreatedAt   time.Time      `json:"-"`
// 	UpdatedAt   time.Time      `json:"-"`
// 	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
// 	Address     Address        `gorm:"foreignKey:AddressID" json:"-"`
// 	User        User           `gorm:"foreignKey:RecipientID" json:"-"`
// 	Request     Request        `gorm:"foreignKey:RequestID" json:"-"`
// }
