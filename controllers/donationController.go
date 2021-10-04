package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"fmt"
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
	var newDonation models.NewDonation
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	c.Bind(&newDonation)
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
// Input : recipient_id, request_id, amount
func UpdateDonationInCartController(c echo.Context) error {
	var update models.CartItemInputData
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	c.Bind(&update)
	update.DonorID = uint(userId)

	if _, err := libdb.GetDonationInCart(update); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to find item in cart",
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
		Message: "success add donation to cart",
		Data:    res,
	})
}

// Delete donation in cart
func DeleteDonationFromCartController(c echo.Context) error {
	var update models.CartItemInputData
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	c.Bind(&update)
	update.DonorID = uint(userId)

	if _, err := libdb.GetDonationInCart(update); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to find item in cart",
		})
	}

	if err := libdb.DeleteDonationFromCart(update); err != nil {
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
// Input : recipient_id, request_id
func CheckoutDonationFromCartController(c echo.Context) error {
	var donation models.DonationCheckout
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)

	c.Bind(&donation)
	donation.DonorID = uint(userId)
	res, err := libdb.CheckoutDonation(donation)
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

// Paid donation status change - unresolved -> resolved donation
// Will make invoice
// Input : donation_id
func PaidDonationController(c echo.Context) error {
	donationId, _ := strconv.ParseUint(c.Param("donation_id"), 0, 0)

	// Find unresolved donation
	tx, err := libdb.GetSpecificDonation(uint(donationId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: fmt.Sprintf("failed to find unresolved donation with id %d", donationId),
		})
	}
	if tx.PaymentStatus == "true" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: fmt.Sprintf("donation with id %d is resolved", donationId),
		})
	}

	paid := "true" // Mock payment process

	res, err := libdb.ChangePaymentStatusToPaid(uint(donationId), paid)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to process payment, fund will be returned",
		})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success make donation",
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
