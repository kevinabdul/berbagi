package controllers

import (
	"net/http"

	libdb "berbagi/lib/database"
	models "berbagi/models"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	var newUser models.RegistrationAPI
	role := c.Param("role")

	newUser.Role = role

	c.Bind(&newUser)

	res, err := libdb.RegisterUser(newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		User models.RegistrationResponseAPI
	}{Status: "success", Message: "Donor has been created!", User: res})

}

// func RegisterVolunteerController(c echo.Context) error {
// 	var newVolunteer models.RegistrationAPI
// 	newDonor.Role = "donor"

// 	c.Bind(&newDonor)

// 	res, err := libdb.RegisterDonor(newDonor)

// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, struct {
// 			Status  string
// 			Message string
// 		}{Status: "Failed", Message: err.Error()})
// 	}

// 	return c.JSON(http.StatusOK, struct {
// 		Status string
// 		Message string
// 		Donor models.DonorAPI
// 	}{Status: "success", Message: "Donor has been created!", Donor: res})
// }