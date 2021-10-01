package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
	"time"
)

// func GetRequestProfileById(userId uint) (interface{}, error) {

// }

func CreateGiftRequest(data models.NewGiftRequest) (models.NewGiftRequestResponseAPI, error) {
	// check package exists & retrieve package name
	var pack models.Package
	if tx := config.Db.First(&pack, data.PackageID); tx.Error != nil {
		return models.NewGiftRequestResponseAPI{}, tx.Error
	} else if tx.RowsAffected == 0 {
		return models.NewGiftRequestResponseAPI{}, errors.New("package doesn't exist")
	}

	var request models.Request
	request.RecipientID = data.UserID
	request.AddressID = data.AddressID
	request.Type = "gift"
	

	if tx := config.Db.Create(&request); tx.Error != nil {
		return models.NewGiftRequestResponseAPI{}, tx.Error
	}
	
	var details models.GiftRequestDetails
	details.RequestID = request.ID
	details.UserID = data.UserID
	details.AddressID = data.AddressID
	details.PackageID = data.PackageID
	details.Quantity = data.Quantity

	if tx := config.Db.Create(&details); tx.Error != nil {
		return models.NewGiftRequestResponseAPI{}, tx.Error
	}

	var res models.NewGiftRequestResponseAPI
	res.RequestID = request.ID
	res.UserID = data.UserID
	res.Quantity = data.Quantity
	res.Package = pack.Name

	return res, nil
}

func CreateDonationRequest(data models.NewDonationRequest) (models.NewDonationRequestResponseAPI, error) {
	var request models.Request
	request.RecipientID = data.FoundationID
	request.AddressID = data.AddressID
	request.Type = "donation"
	

	if tx := config.Db.Create(&request); tx.Error != nil {
		return models.NewDonationRequestResponseAPI{}, tx.Error
	}
	
	var details models.DonationRequestDetails
	details.RequestID = request.ID
	details.UserID = data.FoundationID
	details.AddressID = data.AddressID
	details.Nominal = data.Nominal
	details.Purpose = data.Purpose

	if tx := config.Db.Create(&details); tx.Error != nil {
		return models.NewDonationRequestResponseAPI{}, tx.Error
	}

	var res models.NewDonationRequestResponseAPI
	res.RequestID = request.ID
	res.UserID = data.FoundationID
	res.Nominal = data.Nominal
	res.Purpose = data.Purpose

	return res, nil
}

func CreateServiceRequest(data models.NewServiceRequest) (models.NewServiceRequestResponseAPI, error) {
	// check package exists & retrieve package name
	var serv models.Service
	if tx := config.Db.First(&serv, data.FoundationID); tx.Error != nil {
		return models.NewServiceRequestResponseAPI{}, tx.Error
	} else if tx.RowsAffected == 0 {
		return models.NewServiceRequestResponseAPI{}, errors.New("service doesn't exist")
	}

	var request models.Request
	request.RecipientID = data.FoundationID
	request.AddressID = data.AddressID
	request.Type = "service"
	

	if tx := config.Db.Create(&request); tx.Error != nil {
		return models.NewServiceRequestResponseAPI{}, tx.Error
	}
	
	timeConfig := "2006-01-02"
	var details models.ServiceRequestDetails
	details.RequestID = request.ID
	details.UserID = data.FoundationID
	details.AddressID = data.AddressID
	details.ServiceID = data.ServiceID
	details.StartDate, _ = time.Parse(timeConfig, data.StartDate)
	details.FinishDate, _ = time.Parse(timeConfig, data.FinishDate)

	if tx := config.Db.Create(&details); tx.Error != nil {
		return models.NewServiceRequestResponseAPI{}, tx.Error
	}

	var res models.NewServiceRequestResponseAPI
	res.RequestID = request.ID
	res.UserID = data.FoundationID
	res.Service = serv.Name
	res.StartDate = details.StartDate
	res.FinishDate = details.FinishDate

	return res, nil
}

