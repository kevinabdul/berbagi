package controllers

import (
	"net/http"

	libdb "berbagi/lib/database"
	models "berbagi/models"
	"berbagi/utils/response"

	"github.com/labstack/echo/v4"
)

func RegisterUserController(c echo.Context) error {
	var newUser models.RegistrationAPI

	c.Bind(&newUser)

	res, err := libdb.RegisterUser(newUser)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Create("failed", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Create("success", "user has been created!", res))

}