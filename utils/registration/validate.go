package datavalidation

import (
	"berbagi/models"
	"errors"
	//"reflect"
)

func CheckIncomingData(incomingData *models.RegistrationAPI) error {
	role := incomingData.RoleID
	if !IsRoleValid(role) {
		return errors.New("Invalid role_id. You must choose between 1 (admin), 2 (donor), 3 (volunteer), 4 (children), or 5 (foundation)")
	}

	if incomingData.Email == "" || incomingData.Password == ""{
		return errors.New("Invalid Email or Password. Make sure its not empty and are of string type")
	}

	if incomingData.Latitude == "" || incomingData.Longitude == ""{
		return errors.New("Invalid Latitude or Longitude. Make sure its not empty and are of string type")
	}

	if incomingData.Name == "" {
		return errors.New("Name must be specifed")
	}

	if incomingData.CityID == 0 || incomingData.ProvinceID == 0 {
		return errors.New("CityID and ProvinceID must be specifed")
	}

	if incomingData.NIK == "" {
		return errors.New("Invalid NIK. Make sure its not empty and are of string type")
	}

	if role == 3 && incomingData.ProficiencyID == 0 {
		return errors.New("SkillID must be specifed")
	}

	if role == 5 && incomingData.LicenseID == 0 {
		return errors.New("LicenseID must be specifed")
	}
	return nil
}

func IsRoleValid(role uint) bool {
	roles := []uint{1,2,3,4,5}

	for _, v := range roles {
		if v == role {
			return true
		}
	}
	return false
}