package libdb

import (
	"berbagi/config"
	"berbagi/models"
)

func GetCompletionDetail(verificationId, volunteerId int) (interface{}, int, error) {
	completion := models.Completion{}
	findCompletion := config.Db.Find(&completion, verificationId)
	if findCompletion.Error != nil {
		return nil, 0, findCompletion.Error
	}
	if findCompletion.RowsAffected > 0 {
		confirmedData := models.ConfirmServicesAPI{}
		tx := config.Db.Find(&confirmedData, verificationId)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}

		if confirmedData.VolunteerID != uint(volunteerId) {
			return nil, -1, tx.Error
		}

		response := formattingVerification(confirmedData)
		completionDetail := models.CompletionResponse{
			CompletionStatus: completion.CompletionStatus,
			Invoice:          response.Invoice,
			VolunteerName:    response.VolunteerName,
			AddressVolunteer: response.AddressVolunteer,
			AddressUser:      response.AddressUser,
			ProficiencyName:  response.ProficiencyName,
			UserName:         response.UserName,
			StartDate:        response.StartDate,
			FinishDate:       response.FinishDate,
		}
		return completionDetail, 1, nil
	}
	return nil, 0, nil
}
