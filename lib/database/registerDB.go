package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"berbagi/utils/password"
	datavalidation "berbagi/utils/registration"
	"errors"
	"os"

	"gorm.io/gorm"
)

func RegisterUser(incomingData models.RegistrationAPI) (models.RegistrationResponseAPI, error) {
	dataCheckErr := datavalidation.CheckIncomingData(&incomingData)

	if dataCheckErr != nil {
		return models.RegistrationResponseAPI{}, dataCheckErr
	}

	hashedPassword, err := password.Hash(incomingData.Password)

	if err != nil {
		return models.RegistrationResponseAPI{}, err
	}

	newUser := models.User{}

	transactionErr := config.Db.Transaction(func(tx *gorm.DB) error {

		newAddress := models.Address{}
		newAddress.Name = incomingData.AddressName
		newAddress.Latitude = incomingData.Latitude
		newAddress.Longitude = incomingData.Longitude
		newAddress.CityID = incomingData.CityID
		newAddress.ProvinceID = incomingData.ProvinceID

		if err := tx.Model(models.Address{}).Create(&newAddress).Error; err != nil {
			return err
		}

		newUser.Name = incomingData.Name
		newUser.NIK = incomingData.NIK
		newUser.Email = incomingData.Email
		newUser.Password = hashedPassword
		newUser.RoleID = incomingData.RoleID

		if err := tx.Model(models.User{}).Create(&newUser).Error; err != nil {
			return err
		}

		// Yes, there is a lot of unnecessary duplication of code but i prefer clarity over brevity
		// when doing a non trivial project
		if incomingData.RoleID == 2 {
			newUserRole := models.Donor{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("donors").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 1 {
			// If we want to add credential check when someone register themselves as admin,
			// we could do that here, before adding the new user to admin table
			// e.g: we can define some sort of admin key that must be included in request body
			// the key can be stored in db, cache, or even env file. Then we check for that key
			// every time someone try to register themselves as an admin.
			adminKey := os.Getenv("ADMIN_KEY")

			if adminKey != incomingData.AdminKey || incomingData.AdminKey == "" {
				return errors.New("Invalid admin key")
			}

			newUserRole := models.Admin{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("admins").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 4 {
			newUserRole := models.Children{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("childrens").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 3 {
			newUserRole := models.Volunteer{}
			newUserRole.UserID = newUser.ID
			newUserRole.BirthDate = incomingData.BirthDate
			newUserRole.ProficiencyID = incomingData.ProficiencyID
			newUserRole.AddressID = newAddress.ID

			// addProficiency := tx.Table("proficiencies").Create(&models.Proficiency{ID : incomingData.ProficiencyID})

			// if addProficiency.Error != nil {
			// 	return addProficiency.Error
			// }

			res := tx.Table("volunteers").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		} else if incomingData.RoleID == 5 {
			newUserRole := models.Foundation{}
			newUserRole.UserID = newUser.ID
			newUserRole.LicenseID = incomingData.LicenseID
			newUserRole.AddressID = newAddress.ID

			res := tx.Table("foundations").Create(&newUserRole)

			if res.Error != nil {
				return res.Error
			}
		}

		return nil
	})

	if transactionErr != nil {
		return models.RegistrationResponseAPI{}, transactionErr
	}

	responseAPI := models.RegistrationResponseAPI{}
	responseAPI.UserID = newUser.ID
	responseAPI.Name = newUser.Name
	responseAPI.Email = newUser.Email

	return responseAPI, nil
}
