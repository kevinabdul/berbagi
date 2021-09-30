package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"not null" json:"name" form:"name"`
	Price       int            `gorm:"not null" json:"price" form:"price"`
	CategoryID  uint           `gorm:"not null" json:"category_id" form:"category_id"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
}

type ProductPackage struct {
	ID          uint           `gorm:"primaryKey"`
	Name        string         `gorm:"not null" json:"name" form:"name"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ProductPackageDetail struct {
	ProductPackageID 	uint           	`gorm:"primaryKey" json:"product_package_id" form:"product_package_id"`
	ProductID 			uint           	`gorm:"primaryKey" json:"product_id" form:"product_id"`
	Quantity			int 			`json:"quantity"`
	ProductPackage      ProductPackage
	Product 			Product
}
