package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetListVolunteersController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorized access"})
	}

	volunteer := []models.User{}
	volunteers, rowAffected, err := libdb.ListVolunteers(&volunteer)
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

func GetVolunteerProfileController(c echo.Context) error {
	volunteerId, errorId := strconv.Atoi(c.Request().Header.Get("userId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid volunteer id"})
	}

	volunteer, rowAffected, err := libdb.GetVolunteerProfile(volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get volunteer profile"})
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
	}{Status: "Success", Message: "Success to get volunteer profile", Data: volunteer})
}
