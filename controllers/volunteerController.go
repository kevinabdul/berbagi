package controllers

import (
	libdb "berbagi/lib/database"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetListVolunteers(c echo.Context) error {
	volunteers, rowAffected, err := libdb.ListVolunteers()
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get list volunteers"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "volunteer data not found"})
	}

	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to get list volunteers", Data: volunteers})
}
