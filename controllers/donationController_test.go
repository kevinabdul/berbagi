package controllers_test

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
// e.POST("/donation", handler.MakeDonationController, middlewares.AuthenticateUser)
// e.GET("/cart/donation", handler.GetDonationListInCartController, middlewares.AuthenticateUser)
// e.PUT("/cart/donation", handler.UpdateDonationInCartController, middlewares.AuthenticateUser)
// e.DELETE("/cart/donation", handler.DeleteDonationFromCartController, middlewares.AuthenticateUser)
// e.POST("/donation/checkout", handler.CheckoutDonationFromCartController, middlewares.AuthenticateUser)
// e.PUT("/donation/checkout/:donation_id", handler.PaidDonationController, middlewares.AuthenticateUser)

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
			fmt.Println(body)
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
			recipientId:          "3",
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
			testName:             "success with request",
			path:                 "/donation",
			userId:               "1",
			role:                 "donor",
			recipientId:          "3",
			requestId:            "2",
			amount:               "100000",
			paymentId:            "1",
			quick:                "yes",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "payment",
			expectBodyContains2:  "BERBAGI",
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
