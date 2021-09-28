package models

type Province struct {
	ID 		uint 	`gorm:"primaryKey" json:"id"`
	Name 	string	`gorm:"type:varchar(50)" json:"name"`
}