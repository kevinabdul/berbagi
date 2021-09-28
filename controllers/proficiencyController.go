package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func CreateNewProficiencyController(c echo.Context) error {
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
		}{Status: "Failed", Message: "Failed to create new proficiency"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to create new proficiency", Data: newProficiency})
}

func GetAllProficienciesController(c echo.Context) error {

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
		}{Status: "Failed", Message: "Failed to get all list proficiencies"})
	}
	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "Success to get all list proficiencies", Data: foundProficiency})
}
