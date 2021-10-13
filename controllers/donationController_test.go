package controllers_test

// ACTIVATE TEST SECTION IN LIBDB.PAYMENT.GO

import (
	"berbagi/controllers"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// e.GET("/donation", handler.GetDonationsListController, middlewares.AuthenticateUser)
//- e.POST("/donation", handler.MakeDonationController, middlewares.AuthenticateUser)
//- e.GET("/cart/donation", handler.GetDonationListInCartController, middlewares.AuthenticateUser)
//- e.PUT("/cart/donation", handler.UpdateDonationInCartController, middlewares.AuthenticateUser)
//- e.DELETE("/cart/donation", handler.DeleteDonationFromCartController, middlewares.AuthenticateUser)
//- e.POST("/donation/checkout", handler.CheckoutDonationFromCartController, middlewares.AuthenticateUser)
//x e.PUT("/donation/checkout/:donation_id", handler.PaidDonationController, middlewares.AuthenticateUser)

func TestMakeDonation(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		recipientId          string
		requestId            string
		amount               string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "2",
			amount:               "500000",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "donation",
			expectBodyContains2:  "cart",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"recipient_id": testCase.recipientId,
			"request_id":   testCase.requestId,
			"amount":       testCase.amount,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.MakeDonationController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}
func TestNoRequestDonation(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		recipientId          string
		amount               string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			amount:               "300000",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "donation",
			expectBodyContains2:  "cart",
		},
		{
			testName:             "success",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "4",
			amount:               "400000",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "donation",
			expectBodyContains2:  "request_id\":0",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"recipient_id": testCase.recipientId,
			"amount":       testCase.amount,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.MakeDonationController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestQuickDonation(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		recipientId          string
		requestId            string
		amount               string
		paymentId            string
		quick                string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "failed",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "5",
			amount:               "100000",
			paymentId:            "1",
			quick:                "yes",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "checkout",
			expectBodyContains2:  "donation",
		},
		{
			testName:             "success no request",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "2",
			amount:               "200000",
			paymentId:            "1",
			quick:                "yes",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "payment",
			expectBodyContains2:  "BERBAGI",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"recipient_id": testCase.recipientId,
			"request_id":   testCase.requestId,
			"amount":       testCase.amount,
			"payment_id":   testCase.paymentId,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.QueryParams().Add("quick", testCase.quick)

		if assert.NoError(t, controllers.MakeDonationController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestGetDonationInCart(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success with request",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "cart",
			expectBodyContains2:  "donor",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.GetDonationListInCartController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestUpdateDonationInCart(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		recipientId          string
		requestId            string
		amount               string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "0",
			amount:               "80500",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "cart",
			expectBodyContains2:  "donor",
		},
		{
			testName:             "failed",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "2",
			amount:               "90500",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "only",
			expectBodyContains2:  "non",
		},
		{
			testName:             "failed",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "4",
			requestId:            "2",
			amount:               "80700",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "only",
			expectBodyContains2:  "non",
		},
		{
			testName:             "success with request",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "4",
			requestId:            "0",
			amount:               "90400",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "cart",
			expectBodyContains2:  "donor",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"recipient_id": testCase.recipientId,
			"request_id":   testCase.requestId,
			"amount":       testCase.amount,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPut, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.UpdateDonationInCartController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestDeleteDonationInCart(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		recipientId          string
		requestId            string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "2",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "delete",
			expectBodyContains2:  "cart",
		},
		{
			testName:             "failed",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "3",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "delete",
			expectBodyContains2:  "cart",
		},
		{
			testName:             "success with request",
			path:                 "/cart/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "4",
			requestId:            "0",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "cart",
			expectBodyContains2:  "delete",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"recipient_id": testCase.recipientId,
			"request_id":   testCase.requestId,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodDelete, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.DeleteDonationFromCartController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestCheckoutDonation(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		recipientId          string
		requestId            string
		paymentId            string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/donation/checkout",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "0",
			paymentId:            "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "checkout",
			expectBodyContains2:  "payment",
		},
		{
			testName:             "failed",
			path:                 "/donation/checkout",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "0",
			paymentId:            "1",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "checkout",
			expectBodyContains2:  "donation",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"recipient_id": testCase.recipientId,
			"request_id":   testCase.requestId,
			"payment_id":   testCase.paymentId,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.CheckoutDonationController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestGetDonationResolved(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		resolved             string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			resolved:             "yes",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "[",
		},
		{
			testName:             "failed",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			resolved:             "no",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "[{",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.QueryParams().Add("resolved", testCase.resolved)

		if assert.NoError(t, controllers.GetDonationsListController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestPayDonation(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		invoideId            string
		total                string
		paymentId            string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/payment/donation",
			userId:               "1",
			role:                 "donor",
			invoideId:            "BERBAGI.DONOR.001.DONATE.001.2021-10-13",
			total:                "200000",
			paymentId:            "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "payment success",
			expectBodyContains2:  "",
		},
		{
			testName:             "failed",
			path:                 "/payment/donation",
			userId:               "1",
			role:                 "donor",
			invoideId:            "BERBAGI.VOLUNTEER.001.DONATE.001.2021-10-13",
			total:                "200000",
			paymentId:            "1",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "process",
			expectBodyContains2:  "",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"invoice_id": testCase.invoideId,
			"total":   testCase.total,
			"payment_method_id":   testCase.paymentId,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.AddPendingDonationPaymentController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}
