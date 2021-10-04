package controllers

import (
	"net/http"
	"strconv"

	libdb "berbagi/lib/database"
	models "berbagi/models"
	"berbagi/utils/response"

	"github.com/labstack/echo/v4"
)

func GetPendingPaymentsController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	pendingPayments, err := libdb.GetPendingPaymentsByDonorId(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Create("failed", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Create("success", "Pending payments are retrieved succesfully", pendingPayments))
}

func AddPendingPaymentController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	payment := models.UserPaymentAPI{}
	c.Bind(&payment)
	receiptAPI, err := libdb.AddPendingPaymentByDonorId(payment, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Create("failed", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Create("success", "Payment is succesfull", receiptAPI))
}