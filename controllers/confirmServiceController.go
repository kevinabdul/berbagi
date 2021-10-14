package controllers

import (
	libdb "berbagi/lib/database"
	models "berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddConfirmServiceController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	confirmData, rowAffected, err := libdb.AddConfirmService(volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to confim service activity data"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "volunteer id not found"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to confirm service activity data",
		Data:    confirmData})
}

func GetConfirmServiceController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))
	verificationId, errorId := strconv.Atoi(c.Param("verificationId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid verification id"})
	}

	confirmData, rowAffected, err := libdb.GetConfirmService(verificationId, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get confirmation service data"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "verification id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorize access"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to get confirmation service data",
		Data:    confirmData})
}

func DisplayConfirmServiceController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))
	verificationId, errorId := strconv.Atoi(c.Param("verificationId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid verification id"})
	}

	confirmData, rowAffected, err := libdb.GetConfirmService(verificationId, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get confirmation service data"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "verification id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorize access"})
	}

	return c.Render(http.StatusOK, "letter.html", confirmData)
}
