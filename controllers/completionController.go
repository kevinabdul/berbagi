package controllers

import (
	libdb "berbagi/lib/database"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCompletionDetailController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" && checkRole != "volunteer" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorized access"})
	}

	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	verificationId, errorId := strconv.Atoi(c.Param("verificationId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid verification id"})
	}

	completion, rowAffected, err := libdb.GetCompletionDetail(checkRole, verificationId, volunteerId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get completion data"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "verification id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorize access"})
	}

	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to get completion data", Data: completion})

}

func UpdateStatusCompletionController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorized access"})
	}

	status := c.QueryParam("status")
	if status != "verified" && status != "on-going" && status != "completed" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid completion status"})
	}

	verificationId, errorId := strconv.Atoi(c.Param("verificationId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid verification id"})
	}

	completion, rowAffected, err := libdb.UpdateCompletionStatus(status, verificationId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to update completion status"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "verification id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorize access"})
	}

	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "success to update completion status", Data: completion})

}
