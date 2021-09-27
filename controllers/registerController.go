package controllers

import (
	"net/http"

	libdb "berbagi/lib/database"
	models "berbagi/models"

	"github.com/labstack/echo/v4"
)

func RegisterDonorController(c echo.Context) error {
	var newDonor models.RegistrationAPI
	newDonor.Role = "donor"

	c.Bind(&newDonor)

	res, err := libdb.RegisterDonor(newDonor)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		Donor models.DonorAPI
	}{Status: "success", Message: "Donor has been created!", Donor: res})

}