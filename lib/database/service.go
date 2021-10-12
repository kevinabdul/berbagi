package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"time"
)

func GetServiceByVolunteerId(volunteerId int) (interface{}, int, error) {
	services := models.ServiceCart{}
	tx := config.Db.Where("volunteer_id = ?", volunteerId).Find(&services)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		response := responseService(services)
		return response, 1, nil
	}
	return nil, 0, nil
}

func AddServiceToCart(inputService *models.InputService, volunteerId int) (interface{}, int, error) {
	volunteer := models.Volunteer{}
	tx := config.Db.Find(&volunteer, volunteerId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if tx.RowsAffected > 0 {
		foundation := models.Foundation{}
		findUserId := config.Db.Find(&foundation, inputService.UserID)
		if findUserId.Error != nil {
			return nil, 0, findUserId.Error
		}

		updatedCart := models.ServiceCart{}

		if !ValidatedDate(inputService.StartDate, inputService.FinishDate) {
			return "find another date !", 0, nil
		}

		updatedStartDate := formatDate(inputService.StartDate)
		updatedFinishDate := formatDate(inputService.FinishDate)

		updatedCart.UserID = inputService.UserID
		updatedCart.StartDate = updatedStartDate
		updatedCart.FinishDate = updatedFinishDate
		updatedCart.VolunteerID = uint(volunteerId)
		updatedCart.AddressID = foundation.AddressID

		result := config.Db.Create(&updatedCart)
		if result.Error != nil {
			return nil, 0, result.Error
		}
		return updatedCart, 1, nil
	}

	return nil, 0, nil
}

func DeleteServiceCart(volunteerId int) (interface{}, int, error) {
	serviceCart := models.ServiceCart{}
	tx := config.Db.Delete(&serviceCart, volunteerId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	if tx.RowsAffected > 0 {
		return "deleted", 1, nil
	}
	return nil, 0, nil
}

func UpdatedServiceOnCart(updatedInput *models.InputService, volunteerId int) (interface{}, int, error) {
	service := models.ServiceCart{}
	findService := config.Db.Where("service_carts.volunteer_id = ?", volunteerId).Find(&service)
	if findService.Error != nil {
		return nil, 0, findService.Error
	}

	if findService.RowsAffected > 0 {
		fondation := models.Foundation{}
		findUserId := config.Db.Where("foundations.user_id = ?", updatedInput.UserID).Find(&fondation)
		if findUserId.Error != nil {
			return nil, 0, findUserId.Error
		}

		updatedCart := models.ServiceCart{}
		if !ValidatedDate(updatedInput.StartDate, updatedInput.FinishDate) {
			return "find another date !", 0, nil
		}

		updatedCart.UserID = updatedInput.UserID
		updatedCart.StartDate = formatDate(updatedInput.StartDate)
		updatedCart.FinishDate = formatDate(updatedInput.FinishDate)

		tx := config.Db.Where("service_carts.volunteer_id = ?", volunteerId).Model(&service).Updates(updatedCart)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}
		return service, 1, nil
	}
	return nil, 0, nil
}

func formatDate(date string) time.Time {
	formatedDate, _ := time.Parse("2006-01-02", date)
	return formatedDate
}

func ValidatedDate(date1, date2 string) bool {
	today := time.Now()
	if !today.Before(formatDate(date1)) && !today.Before(formatDate(date2)) {
		return false
	}

	if !formatDate(date1).Before(formatDate(date2)) {
		return false
	}

	MinimRegisterDate := time.Now().Add(7 * 24 * time.Hour)
	if !formatDate(date1).After(MinimRegisterDate) {
		return false
	}

	minimDuration := formatDate(date1).Add(7 * 24 * time.Hour)
	if !formatDate(date2).After(minimDuration) {
		return false
	}

	return true
}

func responseService(serviceCart models.ServiceCart) models.ResponseService {
	volunteer := models.User{}
	findVolunteer := config.Db.Find(&volunteer, serviceCart.VolunteerID)
	if findVolunteer.Error != nil {
		return models.ResponseService{}
	}

	recipient := models.User{}
	findRecipient := config.Db.Find(&recipient, serviceCart.UserID)
	if findRecipient.Error != nil {
		return models.ResponseService{}
	}
	foundation := models.Foundation{}
	findFoundation := config.Db.Find(&foundation, recipient.ID)
	if findFoundation.Error != nil {
		return models.ResponseService{}
	}

	addressName := models.Address{}
	findAddress := config.Db.Find(&addressName, foundation.AddressID)
	if findAddress.Error != nil {
		return models.ResponseService{}
	}

	response := models.ResponseService{
		VolunteerName: volunteer.Name,
		UserName:      recipient.Name,
		AddressName:   addressName.Name,
		StartDate:     serviceCart.StartDate,
		FinishDate:    serviceCart.FinishDate,
	}
	return response
}
