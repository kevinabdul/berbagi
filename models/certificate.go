package models

import (
	"time"

	"gorm.io/gorm"
)

type Certificate struct {
	ID          uint           `gorm:"primaryKey"`
	InvoiceID   uint           `gorm:"not null" json:"invoice_id" form:"invoice_id"`
	Description string         `json:"description" form:"description"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}
