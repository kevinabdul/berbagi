package models

import (
	"time"

	"gorm.io/gorm"
)

type PaymentMethod struct {
	ID        	uint           `gorm:"primaryKey" json:"id"`
	Name      	string         `gorm:"type:varchar(50)" json:"name"`
	Description string 			`gorm:"type:varchar(1000)" json:"description"`
	CreatedAt 	time.Time      `json:"-"`
	UpdatedAt 	time.Time      `json:"-"`
	DeletedAt 	gorm.DeletedAt `gorm:"index"`
}

type PaymentOption struct {
	ID		uint 	`gorm:"primaryKey" json:"payment_method_id"`
	Name	string	`gorm:"unique;type:varchar(25)" json:"payment_method_name"`
}
