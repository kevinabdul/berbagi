package models

import (
	"time"

	"gorm.io/gorm"
)

type Completion struct {
	ConfirmServicesAPIID uint               `gorm:"primaryKey"`
	CompletionStatus     string             `json:"completion_status" sql:"type:ENUM('not verified', 'verified', 'on-going', 'completed')" gorm:"default:not verified"`
	CreatedAt            time.Time          `json:"-"`
	UpdatedAt            time.Time          `json:"-"`
	DeletedAt            gorm.DeletedAt     `gorm:"index" json:"-"`
	ConfirmServicesAPI   ConfirmServicesAPI `gorm:"foreignKey:ConfirmServicesAPIID"`
}

type CompletionResponse struct {
	Invoice          string    `json:"invoice"`
	VolunteerName    string    `json:"volunteer_name"`
	AddressVolunteer string    `json:"volunteer_address"`
	ProficiencyName  string    `json:"proficiency_name"`
	UserName         string    `json:"recipient_name"`
	AddressUser      string    `json:"recipient_address" `
	StartDate        time.Time `json:"start_date" `
	FinishDate       time.Time `json:"finish_date"`
	CompletionStatus string    `json:"completion_status"`
}
