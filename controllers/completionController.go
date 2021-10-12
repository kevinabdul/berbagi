package controllers

import (
	libdb "berbagi/lib/database"
	models "berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCompletionDetailController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" && checkRole != "volunteer" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorized access"})
	}

	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	verificationId, errorId := strconv.Atoi(c.Param("verificationId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid verification id"})
	}

	completion, rowAffected, err := libdb.GetCompletionDetail(checkRole, verificationId, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get completion data"})
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
		Message: "success to get completion data",
		Data:    completion})

}

func UpdateStatusCompletionController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorized access"})
	}

	status := c.QueryParam("status")
	if status != "verified" && status != "on-going" && status != "completed" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid completion status"})
	}

	verificationId, errorId := strconv.Atoi(c.Param("verificationId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid verification id"})
	}

	completion, rowAffected, err := libdb.UpdateCompletionStatus(status, verificationId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to update completion status"})
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
		Message: "success to update completion status",
		Data:    completion})

}
