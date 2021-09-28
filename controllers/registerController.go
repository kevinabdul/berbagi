package controllers

import (
	"berbagi/lib/database"
	"berbagi/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	success = map[string]interface{}{
		"status": "success",
	}
	failed = map[string]interface{}{
		"status": "failed",
	}
	// Reserved for emailUsed case
	emailUsed = map[string]interface{}{
		"status": "failed",
	}
)

func RegisterPersonalRecipientController(c echo.Context) error {
	var recipient models.PersonalRecipients
	c.Bind(&recipient)

	if libdb.IsEmailAvailable(&recipient) == false {
		return c.JSON(http.StatusBadRequest, emailUsed)
	}

	if err := libdb.RegisterPersonalRecipient(&recipient); err != nil {
		return c.JSON(http.StatusBadRequest, failed)
	}

	return c.JSON(http.StatusOK, success)
}

func RegisterAgencyRecipientController(c echo.Context) error {
	var recipient models.AgencyRecipients
	c.Bind(recipient)

	if libdb.IsEmailAvailable(&recipient) == false {
		return c.JSON(http.StatusBadRequest, emailUsed)
	}

	if err := libdb.RegisterAgencyRecipient(&recipient); err != nil {
		return c.JSON(http.StatusBadRequest, failed)
	}
	
	return c.JSON(http.StatusOK, success)
}

func RegisterDonatorController(c echo.Context) error {
	var donator models.Donators
	c.Bind(donator)

	if libdb.IsEmailAvailable(&donator) == false {
		return c.JSON(http.StatusBadRequest, emailUsed)
	}

	if err := libdb.RegisterDonator(&donator); err != nil {
		return c.JSON(http.StatusBadRequest, failed)
	}

	return c.JSON(http.StatusOK, success)
}

func RegisterVolunteerController(c echo.Context) error {
	var volunteer models.Volunteers
	c.Bind(volunteer)

	if libdb.IsEmailAvailable(&volunteer) == false {
		return c.JSON(http.StatusBadRequest, emailUsed)
	}

	if err := libdb.RegisterVolunteer(&volunteer); err != nil {
		return c.JSON(http.StatusBadRequest, failed)
	}

	return c.JSON(http.StatusOK, success)
}