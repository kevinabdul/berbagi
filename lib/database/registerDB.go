package libdb

import (
	"berbagi/config"
	"berbagi/models"
)

const (
	personalRecipientTable = "personalRecipients"
	agencyRecipientTable   = "agencyRecipients"
	donatorTable           = "donators"
	volunteerTable         = "volunteers"
)

func RegisterPersonalRecipient(personalRecipient *models.PersonalRecipients) error {
	if err := config.Db.Table(personalRecipientTable).Create(personalRecipient).Error; err != nil {
		return err
	}
	return nil
}

func RegisterAgencyRecipient(agencyRecipient *models.AgencyRecipients) error {
	if err := config.Db.Table(agencyRecipientTable).Create(agencyRecipient).Error; err != nil {
		return err
	}
	return nil
}

func RegisterDonator(donator *models.Donators) error {
	if err := config.Db.Table(donatorTable).Create(donator).Error; err != nil {
		return err
	}
	return nil
}

func RegisterVolunteer(volunteer *models.Volunteers) error {
	if err := config.Db.Table(volunteerTable).Create(volunteer).Error; err != nil {
		return err
	}
	return nil
}