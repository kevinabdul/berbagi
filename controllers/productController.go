package controllers

import (
	"net/http"
	"strconv"

	libdb "berbagi/lib/database"
	models "berbagi/models"

	"github.com/labstack/echo/v4"
)

func GetProductsController(c echo.Context) error {
	categoryId,_ := strconv.Atoi(c.QueryParam("categoryId"))

	productsTarget, err := libdb.GetProducts(categoryId)

	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status 	string
			Message string
		}{Status: "failed", Message: err.Error()})
	}

	return c.JSON(http.StatusOK, struct {
		Status 		string
		Message 	string	
		Products 	[]models.ProductAPI
	}{Status: "success", Message: "Products retrieval are succesfull", Products: productsTarget})
}
