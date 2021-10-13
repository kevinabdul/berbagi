package models

import (
	"time"

	"gorm.io/gorm"
)

type Request struct {
	ID        uint           `gorm:"primaryKey"`
	UserID    uint           `json:"user_id"`
	AddressID uint           `json:"address_id"`
	Type      string         `json:"type"`
	Resolved  string         `sql:"type:ENUM('true', 'false')" gorm:"default:false" json:"resolved"`
	User      User           `gorm:"foreignKey:UserID" json:"-"`
	Address   Address        `gorm:"foreignKey:AddressID" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
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
	Request        Request        `gorm:"foreignKey:RequestID" json:"-"`
	User           User           `gorm:"foreignKey:UserID" json:"-"`
	Address        Address        `gorm:"foreignKey:AddressID" json:"-"`
	ProductPackage ProductPackage `gorm:"foreignKey:PackageID" json:"-"`
	CreatedAt      time.Time      `json:"-"`
	UpdatedAt      time.Time      `json:"-"`
	DeletedAt      gorm.DeletedAt `gorm:"index" json:"-"`
}

type NewGiftRequestResponseAPI struct {
	RequestID uint   `json:"request_id"`
	UserID    uint   `json:"user_id"`
	Package   string `json:"package"`
	Quantity  int    `json:"quantity"`
}

type RequestProfile struct {
	RequestId   uint    `json:"request_id"`
	Name        string  `json:"name"`
	Role        string  `json:"role"`
	Type        string  `json:"type"`
	Address     string    `json:"address"`
	City        string  `json:"city"`
	Province    string  `json:"province"`
	Distance    float64 `json:"distance"`
}

type NewDonationRequest struct {
	RequestID    uint    `json:"request_id" form:"request_id"`
	FoundationID uint    `json:"user_id" form:"user_id"`
	AddressID    uint    `json:"address_id" form:"address_id"`
	Amount       float64 `json:"amount" form:"amount"`
	Purpose      string  `json:"purpose" form:"purpose"`
}

type DonationRequestDetails struct {
	RequestID uint           `json:"request_id"`
	UserID    uint           `json:"user_id"`
	AddressID uint           `json:"address_id"`
	Amount    float64        `json:"amount" form:"amount"`
	Purpose   string         `json:"purpose" form:"purpose"`
	Request   Request        `gorm:"foreignKey:RequestID" json:"-"`
	User      User           `gorm:"foreignKey:UserID" json:"-"`
	Address   Address        `gorm:"foreignKey:AddressID" json:"-"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type NewDonationRequestResponseAPI struct {
	RequestID uint    `json:"request_id"`
	UserID    uint    `json:"user_id"`
	Amount    float64 `json:"amount" form:"amount"`
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
	RequestID     uint           `json:"request_id"`
	UserID        uint           `json:"user_id"`
	AddressID     uint           `json:"address_id"`
	ProficiencyID uint           `json:"proficiency_id"`
	StartDate     time.Time      `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate    time.Time      `gorm:"not null" json:"finish_date" form:"finish_date"`
	Request       Request        `gorm:"foreignKey:RequestID" json:"-"`
	User          User           `gorm:"foreignKey:UserID" json:"-"`
	Address       Address        `gorm:"foreignKey:AddressID" json:"-"`
	Proficiency   Proficiency    `gorm:"foreignKey:ProficiencyID" json:"-"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
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
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
