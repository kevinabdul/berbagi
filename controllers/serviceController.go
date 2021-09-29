package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddServiceToCartController(c echo.Context) error {
	volunteerId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid volunteer id"})
	}

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
	volunteerId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid volunteer id"})
	}

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
