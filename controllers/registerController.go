package controllers

import (
	"net/http"

	libdb "berbagi/lib/database"
	models "berbagi/models"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	var newUser models.RegistrationAPI

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
	}{Status: "success", Message: "User has been created!", User: res})

}