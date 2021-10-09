package models

import (
	"time"

	"gorm.io/gorm"
)

type Certificate struct {
	CompletionID uint           `gorm:"primaryKey" json:"completion_id"`
	CreatedAt    time.Time      `json:"-"`
	UpdatedAt    time.Time      `json:"-"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

type CertificateResponse struct {
	Invoice         string `json:"certificate_number"`
	VolunteerName   string `json:"volunteer_name"`
	ProficiencyName string `json:"proficiency_name"`
	UserName        string `json:"recipient_name"`
	StartDate       string `json:"start_date"`
	FinishDate      string `json:"finish_date"`
}
