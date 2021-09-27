package models

type Address struct {
	ID 			uint 		`gorm:"primaryKey" json:"id"`
	Name		string 		`gorm:"type:varchar(150)" json:"name"`
	Latitude	string		`gorm:"type:varchar(150)" json:"latitude"`
	Longitude	string		`gorm:"type:varchar(150)" json:"longitude"`
	CityID 		uint 		`json:"city_id"`
	ProvinceID 	uint  		`json:"province_id"`
	City 		City 		`gorm:"foreignKey:CityID"`
	Province  	Province  	`gorm:"foreignKey:ProvinceID"`
}