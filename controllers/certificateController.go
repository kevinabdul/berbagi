package controllers

import (
	libdb "berbagi/lib/database"
	models "berbagi/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCertificateController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	completionId, errorId := strconv.Atoi(c.Param("completionId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid completion id"})
	}

	certificate, rowAffected, err := libdb.GetCertificateService(volunteerId, completionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get certificate"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "completion id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorize access"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success to get certificate",
		Data:    certificate})

}

func CertificateDisplayController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	completionId, errorId := strconv.Atoi(c.Param("completionId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid completion id"})
	}

	certificate, rowAffected, err := libdb.GetCertificateService(volunteerId, completionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get certificate"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, models.ResponseOK{
			Status:  "success",
			Message: "completion id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "unauthorize access"})
	}

	return c.Render(http.StatusOK, "index.html", certificate)
}
