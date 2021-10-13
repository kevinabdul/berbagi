package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// Post donation - donor -> send to cart
// May make donation by request/not
// If not, it's proposed to make donation from foundation profile
// with RequestID=0
// Input : recipient_id, request_id, amount
func MakeDonationController(c echo.Context) error {
	var newDonation models.DonationInputData
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	quick := c.QueryParams().Get("quick")
	if quick == "yes" {
		c.Request().Header.Add("quick", "yes")
		return CheckoutDonationController(c)
	}

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		rcp, _ := strconv.ParseUint(json_map["recipient_id"].(string), 0, 0)
		amt, _ := strconv.ParseInt(json_map["amount"].(string), 0, 0)
		newDonation.RecipientID = uint(rcp)
		newDonation.Amount = int(amt)
		if _, ok := json_map["request_id"]; ok {
			rqs, _ := strconv.ParseUint(json_map["request_id"].(string), 0, 0)
			if rqs > 0 {
				newDonation.RequestID = uint(rqs)
			}
		}
	} else {
		c.Bind(&newDonation)
	}

	newDonation.DonorID = uint(userId)

	res, err := libdb.MakeDonationToCart(newDonation)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to make donation",
		})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success add donation to cart",
		Data:    res,
	})
}

// List donation in cart
func GetDonationListInCartController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	res, err := libdb.ListDonationInCart(uint(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get cart",
		})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success get donations in cart",
		Data:    res,
	})
}

// Update donation in cart
// Only update non-requested donation
// Input : recipient_id, request_id, amount
func UpdateDonationInCartController(c echo.Context) error {
	var update models.DonationInputData
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		rcp, _ := strconv.ParseUint(json_map["recipient_id"].(string), 0, 0)
		amt, _ := strconv.ParseInt(json_map["amount"].(string), 0, 0)
		update.RecipientID = uint(rcp)
		update.Amount = int(amt)
		if _, ok := json_map["request_id"]; ok {
			rqs, _ := strconv.ParseUint(json_map["request_id"].(string), 0, 0)
			if rqs > 0 {
				update.RequestID = uint(rqs)
			}
		}
	} else {
		c.Bind(&update)
	}

	update.DonorID = uint(userId)

	if _, err := libdb.GetDonationInCart(update); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to find item in cart",
		})
	}

	if update.RequestID != 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "only non-requested donation can be updated",
		})
	}

	if err := libdb.UpdateDonationInCart(update); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to update cart",
		})
	}

	res, _ := libdb.GetDonationInCart(update)
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success update donation in cart",
		Data:    res,
	})
}

// Delete donation in cart
func DeleteDonationFromCartController(c echo.Context) error {
	var update models.DonationInputData
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		rcp, _ := strconv.ParseUint(json_map["recipient_id"].(string), 0, 0)
		update.RecipientID = uint(rcp)

		if _, ok := json_map["request_id"]; ok {
			rqs, _ := strconv.ParseUint(json_map["request_id"].(string), 0, 0)
			if rqs > 0 {
				update.RequestID = uint(rqs)
			}
		}
	} else {
		c.Bind(&update)
	}

	update.DonorID = uint(userId)

	if _, err := libdb.GetDonationInCart(update); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to find item in cart",
		})
	}

	row, err := libdb.DeleteDonationFromCart(update)
	if err != nil || row == 0 {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to delete donation from cart",
		})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success delete donation from cart",
	})
}

// Checkout donation from cart to unresolved donation
// May add countdown timer
// Only checkout single donation for now
// Only exact amount of chosen single donation in cart can be checked-out and paid,
// 		if not quick donation
// Only exact-requested-amount to/can be donated, if request_id stated
// Input : recipient_id, request_id
func CheckoutDonationController(c echo.Context) error {
	var donation models.DonationInputData
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	quick := c.Request().Header.Get("quick")

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		rcp, _ := strconv.ParseUint(json_map["recipient_id"].(string), 0, 0)
		pay, _ := strconv.ParseUint(json_map["payment_id"].(string), 0, 0)

		if _, ok := json_map["amount"]; ok {
			amt, _ := strconv.ParseInt(json_map["amount"].(string), 0, 0)
			if amt > 0 {
				donation.Amount = int(amt)
			}
		}

		donation.RecipientID = uint(rcp)
		donation.PaymentID = int(pay)

		if _, ok := json_map["request_id"]; ok {
			rqs, _ := strconv.ParseUint(json_map["request_id"].(string), 0, 0)
			if rqs > 0 {
				donation.RequestID = uint(rqs)
			}
		}
	} else {
		c.Bind(&donation)
	}

	donation.DonorID = uint(userId)

	res, err := libdb.CheckoutDonation(donation, quick)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to checkout donation",
		})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success checkout donation, proceed to payment",
		Data:    res,
	})
}

// View donations list - foundation
func GetDonationsListController(c echo.Context) error {
	resolved := c.QueryParams().Get("resolved")
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	res, err := libdb.GetBulkDonations(uint(userId), resolved)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to get donations list",
		})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success get donations list",
		Data:    res,
	})
}
