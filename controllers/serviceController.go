package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetServiceOnCartController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	services, rowAffected, err := libdb.GetServiceByVolunteerId(volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get service cart"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "service cart not found !"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to get service cart",
		Data:    services})
}

func AddServiceToCartController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	service := models.InputService{}
	c.Bind(&service)

	cart, rowAffected, err := libdb.AddServiceToCart(&service, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to add service to cart"})
	}

	if cart == "find another date !" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "find another date!"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "volunteer id not found"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to add service on cart",
		Data:    cart})
}

func DeleteServiceCartController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	_, rowAffected, err := libdb.DeleteServiceCart(volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to delete service cart"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "volunteer id not found"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to delete service on cart"})
}

func UpdatedServiceOncartController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	updatedInput := models.InputService{}
	c.Bind(&updatedInput)
	updatedService, rowAffected, err := libdb.UpdatedServiceOnCart(&updatedInput, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to update service cart"})
	}

	if updatedService == "find another date !" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "find another date!"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "volunteer id not found"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to update service on carts",
		Data:    updatedService})
}
