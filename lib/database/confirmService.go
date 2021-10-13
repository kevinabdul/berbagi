package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"fmt"
)

func AddConfirmService(volunteerId int) (interface{}, int, error) {
	volunteer := models.Volunteer{}
	findVolunteer := config.Db.Find(&volunteer, volunteerId)
	if findVolunteer.Error != nil {
		return nil, 0, findVolunteer.Error
	}

	serviceCart := models.ServiceCart{}
	query := `SELECT * FROM service_carts
				WHERE service_carts.volunteer_id = ?;`
	tx := config.Db.Raw(query, volunteerId).Scan(&serviceCart)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	confirmData := models.ConfirmServicesAPI{
		VolunteerID: serviceCart.VolunteerID,
		UserID:      serviceCart.UserID,
		StartDate:   serviceCart.StartDate,
		FinishDate:  serviceCart.FinishDate,
	}

	addConfirm := config.Db.Create(&confirmData)
	if addConfirm.Error != nil {
		return nil, 0, addConfirm.Error
	}

	invoice := fmt.Sprintf("%03v/BERBAGI/VOLUNTEER/%03v/%03v", confirmData.ID, volunteer.ProficiencyID, volunteerId)
	confirmData.Invoice = invoice

	saveConfirmData := config.Db.Save(&confirmData)
	if saveConfirmData.Error != nil {
		return nil, 0, saveConfirmData.Error
	}

	if saveConfirmData.RowsAffected > 0 {
		completion := models.Completion{
			ConfirmServicesAPIID: confirmData.ID,
		}
		saveCompletion := config.Db.Create(&completion)
		if saveCompletion.Error != nil {
			return "completion can't create", 0, saveCompletion.Error
		}

		confirmation := formattingVerification(confirmData)
		deletedCart := config.Db.Where("volunteer_id = ?", volunteerId).Delete(&serviceCart)
		if deletedCart.Error != nil {
			return "can't delete cart", 0, tx.Error
		}

		return confirmation, 1, nil
	}

	return nil, 0, nil
}

func GetConfirmService(verificationId, volunteerId int) (interface{}, int, error) {
	verifiedService := models.ConfirmServicesAPI{}
	tx := config.Db.Find(&verifiedService, verificationId)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}

	if verifiedService.VolunteerID != uint(volunteerId) {
		return nil, -1, tx.Error
	}

	if tx.RowsAffected > 0 {
		response := formattingVerification(verifiedService)
		output := models.ResponseVerification{
			Invoice:          response.Invoice,
			VolunteerName:    response.VolunteerName,
			AddressVolunteer: response.AddressVolunteer,
			UserName:         response.UserName,
			ProficiencyName:  response.ProficiencyName,
			AddressUser:      response.AddressUser,
			StartDate:        response.StartDate.Format(layoutUS),
			FinishDate:       response.FinishDate.Format(layoutUS),
		}
		return output, 1, nil
	}
	return nil, 0, nil
}

func formattingVerification(confirmData models.ConfirmServicesAPI) models.ResponseConfirmServices {
	response := models.ResponseConfirmServices{
		Invoice:    confirmData.Invoice,
		StartDate:  confirmData.StartDate,
		FinishDate: confirmData.FinishDate,
	}

	volunteer := models.Volunteer{}
	tx := config.Db.Find(&volunteer, confirmData.VolunteerID)
	if tx.Error != nil {
		return models.ResponseConfirmServices{}
	}

	proficiency := models.Proficiency{}
	tx = config.Db.Find(&proficiency, volunteer.ProficiencyID)
	if tx.Error != nil {
		return models.ResponseConfirmServices{}
	}

	AddressVolunteer := models.Address{}
	tx = config.Db.Find(&AddressVolunteer, volunteer.AddressID)
	if tx.Error != nil {
		return models.ResponseConfirmServices{}
	}
	recipient := models.User{}
	tx = config.Db.Find(&recipient, confirmData.UserID)
	if tx.Error != nil {
		return models.ResponseConfirmServices{}
	}
	AddressRecipient := models.Address{}
	tx = config.Db.Find(&AddressRecipient, recipient.ID)
	if tx.Error != nil {
		return models.ResponseConfirmServices{}
	}
	volunteerName := models.User{}
	tx = config.Db.Find(&volunteerName, confirmData.VolunteerID)
	if tx.Error != nil {
		return models.ResponseConfirmServices{}
	}

	response.VolunteerName = volunteerName.Name
	response.UserName = recipient.Name
	response.AddressVolunteer = AddressVolunteer.Name
	response.AddressUser = AddressRecipient.Name
	response.ProficiencyName = proficiency.Name

	return response
}
