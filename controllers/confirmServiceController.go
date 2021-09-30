package controllers

import (
	libdb "berbagi/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddConfirmServiceController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	confirmData, rowAffected, err := libdb.AddConfirmService(volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to confim service activity data"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "volunteer id not found"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to confirm service activity data", Data: confirmData})
}
