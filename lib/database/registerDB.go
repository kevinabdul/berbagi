package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"berbagi/utils/password"
	"berbagi/utils/registration"
	//"errors"
	"gorm.io/gorm"
)

func RegisterDonor(incomingData models.RegistrationAPI) (models.DonorAPI, error) {
	dataCheckErr := datavalidation.CheckIncomingData(&incomingData)

	if dataCheckErr != nil {
		return models.DonorAPI{}, dataCheckErr
	}
	
	hashedPassword, err := password.Hash(incomingData.Password)

	if err != nil {
		return models.DonorAPI{}, err
	}

	newDonor := models.Donor{}

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

		newDonor.Name = incomingData.Name
		newDonor.Email = incomingData.Email
		newDonor.Password = hashedPassword
		newDonor.NIK = incomingData.NIK
		newDonor.TanggalLahir = incomingData.TanggalLahir
		newDonor.AddressID = newAddress.ID

		// table user
		// insert user dengan email dan password, lat, longitude
		// dapat userId
		// inser addres table
		// dapat address id

		// insert donor table
		// email: kevinabdul@gmail.com password: 1234
		// users: ID, email, password
		// donors: userId, NIK, tanggallahir, addressID
		res := tx.Table("donors").Create(&newDonor)

		if res.Error != nil {
			return res.Error
		}

		return nil
	})

	if transactionErr != nil {
		return models.DonorAPI{}, transactionErr
	}
	
	DonorAPI := models.DonorAPI{}
	DonorAPI.ID = newDonor.ID
	DonorAPI.Name = newDonor.Name
	DonorAPI.Email = newDonor.Email
	
	return DonorAPI, nil
}