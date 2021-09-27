package models

type Donor struct {
	ID 				uint	`gorm:"primaryKey"`
	Name 			string	`gorm:"type:varchar(75)" json:"name"`
	Email 			string	`gorm:"unique; type:varchar(50)" json:"email"`
	Password 		string	`json:"password"`
	NIK 			string	`gorm:"unique type:varchar(16)" json:"nik"`
	TanggalLahir	string	`json:"tanggal_lahir"`
	AddressID 		uint 	`json:"address_id"`
	Address 		Address `gorm:"foreignKey:AddressID"`
}

type RegisterDonorAPI struct {
	Name 			string	`json:"name"`
	Email 			string	`json:"email"`
	Password 		string	`json:"password"`
	NIK 			string	`json:"nik"`
	TanggalLahir	string	`json:"tanggal_lahir"`
	AddressName  	string 	`json:"address_name"`
	Latitude  		string	`json:"lat"`
	Longitude  		string  `json:"long"`
	CityID			uint  	`json:"city_id"`
	ProvinceID  	uint  	`json:"province_id"`
}

type DonorAPI struct {
	ID  			uint  	`json:"id"`
	Name   			string 	`json:"name"`
	Email  			string 	`json:"email"`
}
