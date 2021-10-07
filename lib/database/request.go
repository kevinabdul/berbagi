package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// func GetRequestProfileById(userId uint) (interface{}, error) {

// }

func CreateGiftRequest(data models.NewGiftRequest) (models.NewGiftRequestResponseAPI, error) {
	// check package exists & retrieve package name
	var pack models.ProductPackage
	if tx := config.Db.First(&pack, data.PackageID); tx.Error != nil {
		return models.NewGiftRequestResponseAPI{}, errors.New("package doesn't exist")
	}

	var request models.Request
	request.UserID = data.UserID
	request.AddressID = data.AddressID
	request.Type = "gift"
	request.Resolved = "false"

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
	request.UserID = data.FoundationID
	request.AddressID = data.AddressID
	request.Type = "donation"

	if tx := config.Db.Create(&request); tx.Error != nil {
		return models.NewDonationRequestResponseAPI{}, tx.Error
	}

	var details models.DonationRequestDetails
	details.RequestID = request.ID
	details.UserID = data.FoundationID
	details.AddressID = data.AddressID
	details.Amount = data.Amount
	details.Purpose = data.Purpose

	if tx := config.Db.Create(&details); tx.Error != nil {
		return models.NewDonationRequestResponseAPI{}, tx.Error
	}

	var res models.NewDonationRequestResponseAPI
	res.RequestID = request.ID
	res.UserID = data.FoundationID
	res.Amount = data.Amount
	res.Purpose = data.Purpose

	return res, nil
}

func CreateServiceRequest(data models.NewServiceRequest) (models.NewServiceRequestResponseAPI, error) {
	// check package exists & retrieve package name
	var serv models.Proficiency
	if tx := config.Db.First(&serv, data.ServiceID); tx.Error != nil {
		return models.NewServiceRequestResponseAPI{}, errors.New("service doesn't exist")
	}

	var request models.Request
	request.UserID = data.FoundationID
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
	details.ProficiencyID = data.ServiceID
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

func GetBulkRequests(userId uint, resolved string) ([]models.Request, error) {
	var request []models.Request
	table := "requests"

	if resolved == "yes" {
		tx := config.Db.Table(table).Where("user_id = ? AND resolved = true", userId).Find(&request)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else if resolved == "no" {
		tx := config.Db.Table(table).Where("user_id = ? AND resolved = false", userId).Find(&request)
		if tx.Error != nil {
			return nil, tx.Error
		}
	} else {
		tx := config.Db.Table(table).Where("user_id = ?", userId).Find(&request)
		if tx.Error != nil {
			return nil, tx.Error
		}
	}

	return request, nil
}

func GetTypeRequests(userId uint, reqType, resolved string) (interface{}, error) {
	var joinTable string
	requestTable := "requests"

	if reqType == "gift" {
		request := []models.GiftRequestDetails{}
		joinTable = "gift_request_details"
		if resolved != "" {
			resolve := "false"
			if resolved == "yes" {
				resolve = "true"
			}

			tx := config.Db.Joins(
				fmt.Sprintf(`JOIN %s ON %s.id = %s.request_id`,
					requestTable, requestTable, joinTable)).Where(
				fmt.Sprintf("%s.user_id = %d AND %s.resolved = %s",
					requestTable, userId, requestTable, resolve)).Find(&request)

			if tx.Error != nil {
				return nil, tx.Error
			}
		} else {
			tx := config.Db.Joins(
				fmt.Sprintf(`JOIN %s ON %s.id = %s.request_id`,
					requestTable, requestTable, joinTable)).Where(
				fmt.Sprintf("%s.user_id = %d",
					joinTable, userId)).Find(&request)

			if tx.Error != nil {
				return nil, tx.Error
			}
		}
		return request, nil
	} else if reqType == "donation" {
		request := []models.DonationRequestDetails{}
		joinTable = "donation_request_details"
		if resolved != "" {
			resolve := "false"
			if resolved == "yes" {
				resolve = "true"
			}

			tx := config.Db.Joins(
				fmt.Sprintf(`JOIN %s ON %s.id = %s.request_id`,
					requestTable, requestTable, joinTable)).Where(
				fmt.Sprintf("%s.user_id = %d AND %s.resolved = %s",
					requestTable, userId, requestTable, resolve)).Find(&request)

			if tx.Error != nil {
				return nil, tx.Error
			}
		} else {
			tx := config.Db.Joins(
				fmt.Sprintf(`JOIN %s ON %s.id = %s.request_id`,
					requestTable, requestTable, joinTable)).Where(
				fmt.Sprintf("%s.user_id = %d",
					joinTable, userId)).Find(&request)

			if tx.Error != nil {
				return nil, tx.Error
			}
		}
		return request, nil
	} else if reqType == "service" {
		request := []models.ServiceRequestDetails{}
		joinTable = "service_request_details"
		if resolved != "" {
			resolve := "false"
			if resolved == "yes" {
				resolve = "true"
			}

			tx := config.Db.Joins(
				fmt.Sprintf(`JOIN %s ON %s.id = %s.request_id`,
					requestTable, requestTable, joinTable)).Where(
				fmt.Sprintf("%s.user_id = %d AND %s.resolved = %s",
					requestTable, userId, requestTable, resolve)).Find(&request)

			if tx.Error != nil {
				return nil, tx.Error
			}
		} else {
			tx := config.Db.Joins(
				fmt.Sprintf(`JOIN %s ON %s.id = %s.request_id`,
					requestTable, requestTable, joinTable)).Where(
				fmt.Sprintf("%s.user_id = %d",
					joinTable, userId)).Find(&request)

			if tx.Error != nil {
				return nil, tx.Error
			}
		}
		return request, nil
	}
	
	return nil, errors.New("invalid type")
}

func DeleteRequest(requestId uint) error {
	tx := config.Db.Where("id = ?", requestId).Delete(&models.Request{})
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func GetRequestByIdResolve(requestId uint, resolved string) (models.Request, error) {
	var request models.Request

	request.ID = requestId

	var tx *gorm.DB
	if resolved == "no" {
		tx = config.Db.Where("resolved = 'false'").Find(&request)
	} else if resolved == "yes" {
		tx = config.Db.Where("resolved = 'true'").Find(&request)
	} else {
		tx = config.Db.Table("requests").Find(&request)
	}

	if tx.Error != nil {
		return models.Request{}, tx.Error
	}
	return request, nil
}

func GetRequestByRecipientIdResolve(recipientId uint, resolved string) ([]models.Request, error) {
	var request []models.Request

	var tx *gorm.DB
	if resolved == "no" {
		tx = config.Db.Where("user_id = ? AND resolved = false", recipientId).Find(&request)
	} else if resolved == "yes" {
		tx = config.Db.Where("user_id = ? AND resolved = true", recipientId).Find(&request)
	} else {
		tx = config.Db.Where("user_id = ?", recipientId).Find(&request)
	}

	if tx.Error != nil {
		return []models.Request{}, tx.Error
	}
	return request, nil
}
