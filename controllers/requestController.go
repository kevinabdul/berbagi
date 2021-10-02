package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Failed struct {
	status  string
	message string
}

type Success struct {
	status  string
	message string
	data    interface{}
}

func RequestGift(c echo.Context) error {
	var newRequest models.NewGiftRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "children" {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "Your role can't request gift"})
	}

	c.Bind(&newRequest)
	newRequest.UserID = uint(userId)

	res, err := libdb.CreateGiftRequest(newRequest)
	if err == errors.New("package doesn't exist") {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "package doesn't exist"})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "Can't create new request"})
	}

	return c.JSON(http.StatusOK,
		Success{status: "success",
			message: "Request has been submitted!",
			data:    res})
}

func RequestDonation(c echo.Context) error {
	var newRequest models.NewDonationRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "foundation" {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "Your role can't request donation"})
	}

	c.Bind(&newRequest)
	newRequest.FoundationID = uint(userId)

	res, err := libdb.CreateDonationRequest(newRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "Can't create new request"})
	}

	return c.JSON(http.StatusOK,
		Success{status: "success",
			message: "Request has been submitted!", data: res})
}

func RequestService(c echo.Context) error {
	var newRequest models.NewServiceRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "foundation" {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "Your role can't request gift"})
	}

	c.Bind(&newRequest)
	newRequest.FoundationID = uint(userId)

	res, err := libdb.CreateServiceRequest(newRequest)
	if err == errors.New("service doesn't exist") {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "service doesn't exist"})
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest,
			Failed{status: "Failed",
				message: "Can't create new request"})
	}

	return c.JSON(http.StatusOK,
		Success{status: "success",
			message: "Request has been submitted!",
			data:    res})
}
