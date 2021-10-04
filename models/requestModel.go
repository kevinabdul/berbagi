package models

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	ID          uint           `gorm:"primaryKey"`
	RecipientID uint           `json:"user_id"`
	AddressID   uint           `json:"address_id"`
	Type        string         `json:"type"`
	Resolved    bool           `json:"resolved"`
	User        User           `gorm:"foreignKey:RecipientID"`
	Address     Address        `gorm:"foreignKey:AddressID"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type NewGiftRequest struct {
	RequestID uint `json:"request_id" form:"request_id"`
	UserID    uint `json:"user_id" form:"user_id"`
	AddressID uint `json:"address_id" form:"address_id"`
	PackageID uint `json:"package_id" form:"package_id"`
	Quantity  int  `json:"quantity" form:"quantity"`
}

type GiftRequestDetails struct {
	RequestID      uint           `json:"request_id"`
	UserID         uint           `json:"user_id"`
	AddressID      uint           `json:"address_id"`
	PackageID      uint           `json:"package_id"`
	Quantity       int            `json:"quantity"`
	Request        Request        `gorm:"foreignKey:RequestID"`
	User           User           `gorm:"foreignKey:UserID"`
	Address        Address        `gorm:"foreignKey:AddressID"`
	ProductPackage ProductPackage `gorm:"foreignKey:PackageID"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index"`
}

type NewGiftRequestResponseAPI struct {
	RequestID uint   `json:"request_id"`
	UserID    uint   `json:"user_id"`
	Package   string `json:"package"`
	Quantity  int    `json:"quantity"`
}

type RequestProfile struct {
	RequestId   uint    `json:"request_id"`
	RecipientId uint    `json:"recipient_id"`
	Role        string  `json:"role"`
	AddressID   uint    `json:"address_id"`
	Type        string  `json:"type"`
	Distance    float64 `json:"distance"`
	// PackageID   uint
	// Quantity    int
}

type NewDonationRequest struct {
	RequestID    uint    `json:"request_id" form:"request_id"`
	FoundationID uint    `json:"user_id" form:"user_id"`
	AddressID    uint    `json:"address_id" form:"address_id"`
	Nominal      float64 `json:"nominal" form:"nominal"`
	Purpose      string  `json:"purpose" form:"purpose"`
}

type DonationRequestDetails struct {
	RequestID uint           `json:"request_id"`
	UserID    uint           `json:"user_id"`
	AddressID uint           `json:"address_id"`
	Nominal   float64        `json:"nominal" form:"nominal"`
	Purpose   string         `json:"purpose" form:"purpose"`
	Request   Request        `gorm:"foreignKey:RequestID"`
	User      User           `gorm:"foreignKey:UserID"`
	Address   Address        `gorm:"foreignKey:AddressID"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type NewDonationRequestResponseAPI struct {
	RequestID uint    `json:"request_id"`
	UserID    uint    `json:"user_id"`
	Nominal   float64 `json:"nominal" form:"nominal"`
	Purpose   string  `json:"purpose" form:"purpose"`
}

type NewServiceRequest struct {
	RequestID    uint   `json:"request_id" form:"request_id"`
	FoundationID uint   `json:"user_id" form:"user_id"`
	AddressID    uint   `json:"address_id" form:"address_id"`
	ServiceID    uint   `json:"service_id" form:"service_id"`
	StartDate    string `json:"start_date" form:"start_date"`
	FinishDate   string `json:"finish_date" form:"finish_date"`
}

type ServiceRequestDetails struct {
	RequestID  uint           `json:"request_id"`
	UserID     uint           `json:"user_id"`
	AddressID  uint           `json:"address_id"`
	ServiceID  uint           `json:"service_id"`
	StartDate  time.Time      `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate time.Time      `gorm:"not null" json:"finish_date" form:"finish_date"`
	Request    Request        `gorm:"foreignKey:RequestID"`
	User       User           `gorm:"foreignKey:UserID"`
	Address    Address        `gorm:"foreignKey:AddressID"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}

type NewServiceRequestResponseAPI struct {
	RequestID  uint      `json:"request_id"`
	UserID     uint      `json:"user_id"`
	Service    string    `json:"service"`
	StartDate  time.Time `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate time.Time `gorm:"not null" json:"finish_date" form:"finish_date"`
}

type Service struct {
	ServiceID uint
	Name      string
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
