package controllers

import (
	"net/http"

	libdb "altastore/lib/database"
	models "altastore/models"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	loggingUser := &models.UserAPI{}
	c.Bind(loggingUser)

	token, err := libdb.LoginUser(loggingUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status string
			Message string
		}{Status: "failed", Message: err.Error()})
	}
	
	return c.JSON(http.StatusOK, struct {
		Status string
		Message string
		Token string
	}{Status: "success", Message: "You are logged in!", Token: token})
}