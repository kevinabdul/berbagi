package models

type RegistrationAPI struct {
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
	SkillID  		uint  	`json:"skill_id"`
	YayasanID  		uint  	`json:"yayasan_id"`
	Role   			string 	`json:"role"`
}