package controllers

import (
	"net/http"

	libdb "berbagi/lib/database"
	models "berbagi/models"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	loggingUser := models.LoginUserAPI{}
	
	c.Bind(&loggingUser)

	token, err := libdb.LoginUser(loggingUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Message string
		}{Message: err.Error()})
	}
	
	return c.JSON(http.StatusOK, models.LoginResponseAPI{Message: "You are logged in!", Token: token})
}