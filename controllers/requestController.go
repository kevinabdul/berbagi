package controllers

import (
	libdb "berbagi/lib/database"
	"berbagi/models"
	"encoding/json"
	"net/http"
	"regexp"
	"strconv"

	"github.com/labstack/echo/v4"
)

func RequestGift(c echo.Context) error {
	var newRequest models.NewGiftRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "children" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't request gift"})
	}

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		addr, _ := strconv.ParseUint(json_map["address_id"].(string), 0, 0)
		pkg, _ := strconv.ParseUint(json_map["package_id"].(string), 0, 0)
		qty, _ := strconv.ParseInt(json_map["quantity"].(string), 0, 0)
		newRequest.AddressID = uint(addr)
		newRequest.PackageID = uint(pkg)
		newRequest.Quantity = int(qty)
	} else {
		c.Bind(&newRequest)
	}

	newRequest.UserID = uint(userId)

	res, err := libdb.CreateGiftRequest(newRequest)
	if err != nil {
		if err.Error() == "package doesn't exist" {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "package doesn't exist"})
		} else {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't create new request"})
		}
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "request has been submitted!",
		Data:    res})
}

func RequestDonation(c echo.Context) error {
	var newRequest models.NewDonationRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't request donation"})
	}

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		addr, _ := strconv.ParseUint(json_map["address_id"].(string), 0, 0)
		amt, _ := strconv.ParseFloat(json_map["amount"].(string), 64)
		pps, _ := json_map["purpose"].(string)
		newRequest.AddressID = uint(addr)
		newRequest.Amount = amt
		newRequest.Purpose = pps
	} else {
		c.Bind(&newRequest)
	}

	newRequest.FoundationID = uint(userId)

	res, err := libdb.CreateDonationRequest(newRequest)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't create new request"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "request has been submitted!",
		Data:    res})
}

func RequestService(c echo.Context) error {
	var newRequest models.NewServiceRequest
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")

	if role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't request service"})
	}

	if c.Request().Header.Get("Content-Type") == "application/json" {
		json_map := make(map[string]interface{})
		err := json.NewDecoder(c.Request().Body).Decode(&json_map)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't parse new request"})
		}

		addr, _ := strconv.ParseUint(json_map["address_id"].(string), 0, 0)
		srv, _ := strconv.ParseUint(json_map["service_id"].(string), 0, 0)
		start, _ := json_map["start_date"].(string)
		finish, _ := json_map["finish_date"].(string)
		newRequest.AddressID = uint(addr)
		newRequest.ServiceID = uint(srv)
		newRequest.StartDate = start
		newRequest.FinishDate = finish
	} else {
		c.Bind(&newRequest)
	}

	newRequest.FoundationID = uint(userId)

	res, err := libdb.CreateServiceRequest(newRequest)
	if err != nil {
		if err.Error() == "service doesn't exist" {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "service doesn't exist"})
		} else {
			return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
				Status:  "failed",
				Message: "can't create new request"})
		}
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "request has been submitted!",
		Data:    res})
}

// Used by children & foundation role
func GetAllRequestListController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")
	resolved := c.QueryParams().Get("resolved")

	if role != "children" && role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't reach this"})
	}

	res, err := libdb.GetBulkRequests(uint(userId), resolved)

	roleStatus, _ := libdb.CheckUserRoleRightness(uint(userId), role)
	if !roleStatus {
		res, err = libdb.GetBulkRequests(uint(userId), "no")
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't get request list"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success getting request list",
		Data:    res})
}

func IdTypeRedirector(c echo.Context) error {
	digitCheck := regexp.MustCompile(`^[0-9]+$`)
	field := c.Param("field")

	if field == "gift" || field == "donation" || field == "service" {
		return GetTypeRequestListController(c)
	} else if digitCheck.MatchString(field) {
		return GetRequestByRecipientIdController(c)
	} else {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid parameter"})
	}
}

func GetTypeRequestListController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	role := c.Request().Header.Get("role")
	reqType := c.Param("field")
	resolved := c.QueryParams().Get("resolved")

	if reqType != "gift" && reqType != "donation" && reqType != "service" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "invalid request type"})
	}

	if role != "children" && role != "foundation" {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't reach this"})
	}

	var res interface{}
	var err error
	if (reqType == "gift" && role == "children") || (reqType == "donation" && role == "foundation") || (reqType == "service" && role == "foundation") {
		// Get gift requests list
		res, err = libdb.GetTypeRequests(uint(userId), reqType, resolved)
	} else {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "your role can't reach this"})
	}

	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't get request list"})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success getting request list",
		Data:    res})
}

func DeleteRequestController(c echo.Context) error {
	userId, _ := strconv.ParseUint(c.Request().Header.Get("userId"), 0, 0)
	requestId, _ := strconv.ParseUint(c.Param("request_id"), 0, 0)

	get, err := libdb.GetRequestByIdResolve(uint(requestId), "no")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't find unresolved request"})
	}

	if uint(userId) != get.UserID {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't delete other's request"})
	}

	if err := libdb.DeleteRequest(uint(requestId)); err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "failed to delete request",
		})
	}

	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success delete request",
	})
}

func GetRequestByRecipientIdController(c echo.Context) error {
	recipientId, _ := strconv.ParseUint(c.Param("field"), 0, 0)

	res, err := libdb.GetRequestByRecipientIdResolve(uint(recipientId), "no")
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.ResponseNotOK{
			Status:  "failed",
			Message: "can't find request list"})
	}
	return c.JSON(http.StatusOK, models.ResponseOK{
		Status:  "success",
		Message: "success getting request list",
		Data:    res,
	})
}
