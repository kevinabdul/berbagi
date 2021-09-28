package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductCart struct {
	DonorID   uint           `gorm:"primaryKey"`
	UserID    uint           `json:"recipient_id" form:"recipient_id"`
	ProductID uint           `gorm:"not null;primaryKey" json:"product_id" form:"product_id"`
	AddressID uint           `gorm:"not null" json:"recipient_address_id" form:"recipient_address_id"`
	Quantity  int            `gorm:"not null" json:"quantity" form:"quantity"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Donor     Donor          `gorm:"foreignKey:DonorID"`
	Product   Product        `gorm:"foreignKey:ProductID"`
	Address   Address        `gorm:"foreignKey:AddressID"`
	User      User           `gorm:"foreignKey:UserID"`
}
