package controllers

import (
	"net/http"
	"strconv"

	libdb "berbagi/lib/database"
	models "berbagi/models"
	"berbagi/utils/response"

	"github.com/labstack/echo/v4"
)

func GetCheckoutByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	checkout, err := libdb.GetCheckoutByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Create("failed", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Create("success", "cart is retrieved succesfully!", checkout))
}

func AddCheckoutByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	payment := models.PaymentMethod{}
	c.Bind(&payment)
	transactionAPI, err := libdb.AddCheckoutByUserId(payment, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Create("failed", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Create("success", "checkout is succesfull", transactionAPI))
}