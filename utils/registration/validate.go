package datavalidation

import (
	"berbagi/models"
	"errors"
	//"reflect"
)

func CheckIncomingData(incomingData *models.RegistrationAPI) error {
	role := incomingData.Role
	if !IsRoleValid(role) {
		return errors.New("Invalid role. You must choose between admin, donor, volunteer, children, or foundation")
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

	if role == "volunteer" && incomingData.ProficiencyID == 0 {
		return errors.New("SkillID must be specifed")
	}

	if role == "foundation" && incomingData.LicenseID == 0 {
		return errors.New("LicenseID must be specifed")
	}
	return nil
}

func IsRoleValid(role string) bool {
	roles := []string{"admin", "donor", "volunteer", "children", "foundation"}

	for _, v := range roles {
		if v == role {
			return true
		}
	}
	return false
}