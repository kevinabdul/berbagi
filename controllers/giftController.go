package controllers

import (
	"net/http"
	"strconv"

	libdb "berbagi/lib/database"
	"berbagi/utils/response"

	"github.com/labstack/echo/v4"
)

func GetGiftsController(c echo.Context) error {
	childrenId , _ := strconv.Atoi(c.Request().Header.Get("userId"))
	status := c.QueryParam("status")

	gifts, err := libdb.GetGiftsByChildrenId(childrenId, status)

	if err != nil {
		return c.JSON(http.StatusBadRequest, response.Create("failed", err.Error(), nil))
	}

	return c.JSON(http.StatusOK, response.Create("success", "gifts are retrieved succesfully", gifts))
}