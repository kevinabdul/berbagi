package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllNearestRecipientsController(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Request().Header.Get("userId"))
	role := c.Request().Header.Get("role")
	_range, _ := strconv.ParseFloat(c.QueryParams().Get("range"), 64)

	if _range > 15 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Range too far"})
	}
	if _range < 0 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Range value invalid"})
	}

	var addressId uint
	var err error
	if role == "donor" {
		addressId, err = libdb.GetAddressIdByUserIdDonor(userId)

	} else if role == "volunteer" {
		addressId, err = libdb.GetAddressIdByUserIdVolunteer(userId)
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get addressID"})
	}

	// var address models.Address
	address, err := libdb.GetAddressById(addressId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get user's address"})
	}
	addr := address.(models.Address)

	lat, _ := strconv.ParseFloat(addr.Latitude, 64)
	lon, _ := strconv.ParseFloat(addr.Longitude, 64)
	nearesRecipientsId, rowAffected, err := libdb.GetAllNearestAddressId(lat, lon, _range)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get nearest recipients"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "No recipients nearby"})
	}

	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success getting recipients nearby", Data: nearesRecipientsId})
}
