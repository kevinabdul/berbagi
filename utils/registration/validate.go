package datavalidation

import (
	"berbagi/models"
	"errors"
	//"reflect"
)

func CheckIncomingData(incomingData *models.RegistrationAPI) error {
	if incomingData.Email == "" || incomingData.Password == ""{
		return errors.New("Invalid Email or Password. Make sure its not empty and are of string type")
	}

	if incomingData.NIK == "" {
		return errors.New("Invalid NIK. Make sure its not empty and are of string type")
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

	if incomingData.Role == "volunteer" && incomingData.ProficiencyID == 0 {
		return errors.New("SkillID must be specifed")
	}

	if incomingData.Role == "foundation" && incomingData.LicenseID == 0 {
		return errors.New("LicenseID must be specifed")
	}
	return nil
}