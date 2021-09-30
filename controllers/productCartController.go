package controllers

import (
	"net/http"
	"strconv"

	libdb "berbagi/lib/database"
	models "berbagi/models"

	"github.com/labstack/echo/v4"
)

func GetProductCartByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	cartTarget, err := libdb.GetProductCartByUserId(userId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message string
		Carts 	[]models.ProductCartGetAPI
	}{Status: "success", Message: "Cart is retrieved succesfully" , Carts: cartTarget})
}


// Update is for addition and update of item(s). Subsequent request without previously added item(s) wont discard the already added item(s)
// Attempt to set quantity if an item to zero will be ignored
// You should use delete endpoint to delete item(s)
func UpdateProductCartByUserIdController(c echo.Context) error {
	donorId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	
	var userCart []models.ProductCart
	c.Bind(&userCart)
	
	err := libdb.UpdateProductCartByUserId(userCart, donorId)	

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message	string
	}{Status: "success", Message: "ProductCart is updated!"})
}

func DeleteProductCartByUserIdController(c echo.Context) error {
	userId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	
	var userCart []models.ProductCartDelAPI
	c.Bind(&userCart)
	
	err := libdb.DeleteProductCartByUserId(userCart, userId)	

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 	string
		Message	string
	}{Status: "success", Message: "Cart is updated!"})
}