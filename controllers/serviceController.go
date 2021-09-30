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
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get service cart"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "service cart not found !"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to get service cart", Data: services})
}

func AddServiceToCartController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	service := models.InputService{}
	c.Bind(&service)

	cart, rowAffected, err := libdb.AddServiceToCart(&service, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to add service to cart"})
	}

	if cart == "find another date !" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Find another date!"})
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
	}{Status: "Success", Message: "Success to add service on cart", Data: cart})
}

func DeleteServiceCartController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	_, rowAffected, err := libdb.DeleteServiceCart(volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to delete service cart"})
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
	}{Status: "Success", Message: "Success to delete service on cart"})
}

func UpdatedServiceOncartController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	updatedInput := models.InputService{}
	c.Bind(&updatedInput)
	updatedService, rowAffected, err := libdb.UpdatedServiceOnCart(&updatedInput, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to update service cart"})
	}

	if updatedService == "find another date !" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Find another date!"})
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
	}{Status: "Success", Message: "success to update service on carts", Data: updatedService})
}
