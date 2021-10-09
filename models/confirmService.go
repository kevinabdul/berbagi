package models

import (
	"time"
)

type ConfirmServicesAPI struct {
	ID          uint      `gorm:"primaryKey"`
	Invoice     string    `json:"invoice"`
	VolunteerID uint      `json:"volunteer_id"`
	UserID      uint      `gorm:"not null" json:"recipient_id" form:"recipient_id"`
	StartDate   time.Time `gorm:"not null" json:"start_date" form:"start_date"`
	FinishDate  time.Time `gorm:"not null" json:"finish_date" form:"finish_date"`
}

type ResponseConfirmServices struct {
	Invoice          string    `json:"invoice"`
	VolunteerName    string    `json:"volunteer_name"`
	AddressVolunteer string    `json:"volunteer_address"`
	ProficiencyName  string    `json:"proficiency_name"`
	UserName         string    `json:"recipient_name"`
	AddressUser      string    `json:"recipient_address" form:"recipient_address"`
	StartDate        time.Time `json:"start_date" form:"start_date"`
	FinishDate       time.Time `json:"finish_date" form:"finish_date"`
}

type ResponseVerification struct {
	Invoice          string `json:"invoice"`
	VolunteerName    string `json:"volunteer_name"`
	AddressVolunteer string `json:"volunteer_address"`
	ProficiencyName  string `json:"proficiency_name"`
	UserName         string `json:"recipient_name"`
	AddressUser      string `json:"recipient_address" form:"recipient_address"`
	StartDate        string `json:"start_date" form:"start_date"`
	FinishDate       string `json:"finish_date" form:"finish_date"`
}
