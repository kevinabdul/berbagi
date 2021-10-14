package controllers_test

import (
	"berbagi/controllers"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestGetNearbyTarget(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		resource             string
		_type                string
		_range               string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "failed - range too far",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "recipient",
			_type:                "foundation",
			_range:               "30",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "range",
			expectBodyContains2:  "far",
		},
		{
			testName:             "failed - range invalid",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "recipient",
			_type:                "foundation",
			_range:               "-3",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "range",
			expectBodyContains2:  "invalid",
		},
		{
			testName:             "failed - invalid resource",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "gift",
			_type:                "foundation",
			_range:               "3",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "resource",
			expectBodyContains2:  "invalid",
		},
		{
			testName:             "failed - invalid category",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "recipient",
			_type:                "volunteer",
			_range:               "3",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "category",
			expectBodyContains2:  "invalid",
		},
		{
			testName:             "failed - mismatch category",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "recipient",
			_type:                "gift",
			_range:               "3",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "combination",
			expectBodyContains2:  "invalid",
		},
		{
			testName:             "success - no target nearby",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "request",
			_type:                "service",
			_range:               "3",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "no",
			expectBodyContains2:  "target",
		},
		{
			testName:             "success - recipient",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "recipient",
			_type:                "children",
			_range:               "3",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "recipient",
			expectBodyContains2:  "nearby",
		},
		{
			testName:             "success - request",
			path:                 "/nearby",
			userId:               "1",
			role:                 "donor",
			resource:             "request",
			_type:                "gift",
			_range:               "3",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "request",
			expectBodyContains2:  "nearby",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.SetParamNames("resource")
		c.SetParamValues(testCase.resource)
		c.QueryParams().Add("type", testCase._type)
		c.QueryParams().Add("range", testCase._range)

		if assert.NoError(t, controllers.GetAllNearestRecipientsController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(testCase.testName,body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}
