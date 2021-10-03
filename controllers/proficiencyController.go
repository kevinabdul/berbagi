package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func CreateNewProficiencyController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorized access"})
	}

	proficiency := models.Proficiency{}
	c.Bind(&proficiency)

	newProficiency, rowAffected, err := libdb.CreateNewProficiency(&proficiency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to create new proficiency"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "New proficiency not found"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to create new proficiency", Data: newProficiency})
}

func GetAllProficienciesController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" && checkRole != "volunteer" && checkRole != "foundation" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorized access"})
	}

	foundProficiency, rowAffected, err := libdb.GetAllProficiencies()
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get all list proficiencies"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "List proficiencies not found"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to get list proficiencies", Data: foundProficiency})
}

func DeleteProficiencyController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorized access"})
	}

	proficiencyId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid proficiency id"})
	}

	deletedMessage, rowAffected, err := libdb.DeleteProficiency(proficiencyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to delete proficiency"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "Proficiency not found"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to delete proficiency", Data: deletedMessage})
}

func UpdatedProficiencyController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorized access"})
	}

	proficiencyId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid proficiency id"})
	}

	newProficiency := models.Proficiency{}
	c.Bind(&newProficiency)

	updatedProficiency, rowAffected, err := libdb.UpdateProficiency(proficiencyId, &newProficiency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to update proficiency"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "Proficiency not found"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to update proficiency", Data: updatedProficiency})
}
