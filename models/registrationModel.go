package models

type RegistrationAPI struct {
	UserID		 	uint   `json:"user_id"`
	Name         	string `json:"name"`
	Email        	string `json:"email"`
	Password     	string `json:"password"`
	NIK          	string `json:"nik"`
	BirthDate	 	string `json:"birth_date"`
	AddressName  	string `json:"address_name"`
	Latitude     	string `json:"latitude"`
	Longitude    	string `json:"longitude"`
	CityID       	uint   `json:"city_id"`
	ProvinceID   	uint   `json:"province_id"`
	ProficiencyID	uint   `json:"proficiency_id"`
	LicenseID    	uint   `json:"license_id"`
	RoleID         	uint `json:"role_id"`
	AdminKey 		string `json:"admin_key"`
}

type RegistrationResponseAPI struct {
	UserID  uint   `json:"user_id"`
	Name  	string `json:"name"`
	Email 	string `json:"email"`
}