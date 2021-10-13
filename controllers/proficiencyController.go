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
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorized access"})
	}

	proficiency := models.Proficiency{}
	c.Bind(&proficiency)

	newProficiency, _, err := libdb.CreateNewProficiency(&proficiency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to create new proficiency"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to create new proficiency",
		Data:    newProficiency})
}

func GetAllProficienciesController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" && checkRole != "volunteer" && checkRole != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorized access"})
	}

	foundProficiency, rowAffected, err := libdb.GetAllProficiencies()
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get all list proficiencies"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "list proficiencies not found"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to get list proficiencies",
		Data:    foundProficiency})
}

func DeleteProficiencyController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorized access"})
	}

	proficiencyId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid proficiency id"})
	}

	deletedMessage, rowAffected, err := libdb.DeleteProficiency(proficiencyId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to delete proficiency"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "proficiency not found"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to delete proficiency",
		Data:    deletedMessage})
}

func UpdatedProficiencyController(c echo.Context) error {
	checkRole := c.Request().Header.Get("role")
	if checkRole != "admin" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorized access"})
	}

	proficiencyId, errorId := strconv.Atoi(c.Param("id"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid proficiency id"})
	}

	newProficiency := models.Proficiency{}
	c.Bind(&newProficiency)

	updatedProficiency, rowAffected, err := libdb.UpdateProficiency(proficiencyId, &newProficiency)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to update proficiency"})
	}
	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "proficiency not found"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to update proficiency",
		Data:    updatedProficiency})
}
