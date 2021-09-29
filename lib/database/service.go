package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"time"
)

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
