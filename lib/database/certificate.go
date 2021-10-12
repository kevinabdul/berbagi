package libdb

import (
	"berbagi/config"
	"berbagi/models"
	"errors"
)

const (
	layoutUS = "January 2, 2006"
)

func GetCertificateService(volunteerId int, completionId int) (interface{}, int, error) {
	completion := models.Completion{}
	findCompletion := config.Db.Find(&completion, completionId)
	if findCompletion.Error != nil {
		return nil, 0, findCompletion.Error
	}

	if findCompletion.RowsAffected > 0 {
		verifiedService := models.ConfirmServicesAPI{}
		tx := config.Db.Find(&verifiedService, completionId)
		if tx.Error != nil {
			return nil, 0, tx.Error
		}

		if verifiedService.VolunteerID != uint(volunteerId) {
			return nil, -1, errors.New("unauthorized access")
		}

		verifiedData := formattingVerification(verifiedService)
		response := models.CertificateResponse{
			Invoice:         verifiedData.Invoice,
			VolunteerName:   verifiedData.VolunteerName,
			UserName:        verifiedData.UserName,
			StartDate:       verifiedData.StartDate.Format(layoutUS),
			FinishDate:      verifiedData.FinishDate.Format(layoutUS),
			ProficiencyName: verifiedData.ProficiencyName,
		}
		return response, 1, nil
	}
	return nil, 0, nil
}
