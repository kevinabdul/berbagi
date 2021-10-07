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
			Status: "failed", 
			Message: "range too far"})
	}
	if _range < 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status: "failed", 
			Message: "range value invalid"})
	}
// ----------------------------------------------------------------------------
	if resource != "request" && resource != "recipient" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status: "failed", 
			Message: "invalid resource type"})
	}
	if _type != "" && _type != "gift" &&
	   _type != "donation" && _type != "service" &&
	   _type != "foundation" && _type != "children" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status: "failed", 
			Message: "invalid category"})
	   }
// ----------------------------------------------------------------------------
	address, err := libdb.GetAddressLatLonByUserId(uint(userId), role)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status: "failed", 
			Message: "failed to get user's address"})
	}
// ----------------------------------------------------------------------------
	lat := address.Latitude
	lon := address.Longitude
	nearestAddressId, rowAffected, err := libdb.GetAllNearestAddressId(lat, lon, _range)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status: "failed", 
			Message: "failed to get nearest recipients"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status: "success", 
			Message: "no target nearby"})
	}
// ----------------------------------------------------------------------------
	var res interface{}
	if resource == "request" {
		if _type == "" || _type == "gift" || _type == "donation" || _type == "service" {
			// get request
			request := []models.RequestProfile{}
			for _, elem := range nearestAddressId {
				res, _, _ := libdb.GetRequestByAddressIdResolve(elem.ID, "no")
				for _, r := range res {
					req := models.RequestProfile{
						RequestId: r.ID,
						RecipientId: r.UserID,
						AddressID: r.AddressID,
						Type: r.Type,
						Distance: elem.Distance,
					}
					request = append(request, req)
				}
			}
			res = request
		} else {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status: "failed", 
				Message: "invalid resource and type combination"})
		}
	} else {
		if _type == "" || _type == "foundation" || _type == "children" {
			// get recipient
			recipient := []models.UserProfile{}
			for _, elem := range nearestAddressId {
				user, _, _ := libdb.GetUserByAddressIdRole(elem.ID, _type)
				recipient = append(recipient, user)
			}
			res = recipient
		} else {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status: "failed", 
				Message: "invalid resource and type combination"})
		}
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status: "success", 
		Message: "success getting target nearby", 
		Data: res})
}
