package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetAllNearestRecipientsController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")
	resource := c.Param("resource")
	_type := c.QueryParams().Get("type")
	_range, _ := strconv.ParseFloat(c.QueryParams().Get("range"), 64)
	// ----------------------------------------------------------------------------
	if _range > 25 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "range too far"})
	}
	if _range < 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "range value invalid"})
	}
	// ----------------------------------------------------------------------------
	if resource != "request" && resource != "recipient" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid resource type"})
	}
	if _type != "" && _type != "gift" &&
		_type != "donation" && _type != "service" &&
		_type != "foundation" && _type != "children" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid category"})
	}
	if !(resource == "recipient" && _type == "foundation" ||
		resource == "recipient" && _type == "children" ||
		resource == "request" && _type == "gift" ||
		resource == "request" && _type == "donation" ||
		resource == "request" && _type == "service") {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid resource and type combination"})
	}
	// ----------------------------------------------------------------------------
	address, err := libdb.GetAddressLatLonByUserId(uint(userId), role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get user's address"})
	}
	// ----------------------------------------------------------------------------
	var getNearby models.NearbyInputData
	getNearby.UserID = uint(userId)
	getNearby.Role = role
	getNearby.Latitude = address.Latitude
	getNearby.Longitude = address.Longitude
	getNearby.Range = _range
	getNearby.GetResource = resource
	getNearby.Type = _type

	// ----------------------------------------------------------------------------
	if resource == "recipient" {
		nearbyUsers, err := libdb.GetAllNearestUsers(getNearby)

		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "failed to get nearest recipients"})
		}

		if len(nearbyUsers) == 0 {
			return c.JSON(http.StatusOK, models.ResponseOK{
				Status:  "success",
				Message: "no target nearby"})
		}

		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "success getting recipients nearby",
			Data:    nearbyUsers})
	}
	// ----------------------------------------------------------------------------
	// if resource == "request"
	request, err := libdb.GetNearbyRequestProfile(getNearby)

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get nearest requests"})
	}

	if len(request) == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "no target nearby"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success getting requests nearby",
		Data:    request})
}
