package controllers

import (
	"encoding/json"
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

	return c.JSON(http.StatusOK, response.Create("success", "pending payments are retrieved succesfully", pendingPayments))
}

func AddPendingPaymentController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	payment := models.UserPaymentAPI{}
	c.Bind(&payment)
	receiptAPI, err := libdb.AddPendingPaymentByDonorId(payment, userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Create("failed", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Create("success", "payment is succesfull", receiptAPI))
}


func AddPendingDonationPaymentController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	payment := models.UserPaymentAPI{}

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		inv, _ := json_map["invoice_id"].(string)
		ttl, _ := strconv.ParseInt(json_map["total"].(string), 0, 0)
		pay, _ := strconv.ParseInt(json_map["payment_method_id"].(string), 0, 0)
		payment.InvoiceID = inv
		payment.Total = int(ttl)
		payment.PaymentMethodID = uint(pay)
	} else {
		c.Bind(&payment)
	}
	// c.Bind(&payment)

	receiptAPI, err := libdb.AddPendingDonationPaymentByDonorId(payment, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status: "failed",
			Message: "failed to process payment; "+err.Error(),
		})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status: "success",
		Message: "payment success",
		Data: receiptAPI,
	})
}