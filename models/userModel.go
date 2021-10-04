package models

import (
	"time"

	"gorm.io/gorm"
)

// Karena login behavior bergantung pada value role query, login controller harus mengekstrak informasi tersebut
// dan melakukan binding ke satu struktur yang uniform seperti dibawah ini
type User struct {
	ID        uint           `gorm:"primaryKey"`
	Name      string         `gorm:"type:varchar(75)" json:"name"`
	Role 	  string		 `gorm:"type:varchar(50)" json:"role"`
	NIK       string         `gorm:"unique; type:varchar(16)" json:"nik"`
	Email     string         `gorm:"unique; type:varchar(50)" json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type LoginUserAPI struct {
	Username string
	Email    string
	Password string
	Role     string
}

type LoginSearchAPI struct {
	ID       uint
	Name     string
	Email    string
	Password string
	Role 	 string
}

type LoginResponseAPI struct {
	Message string
	Token   string
}

type UserProfile struct {
	UserID uint
	Name string
	Role string
	Address string
	City string
	Province string
}