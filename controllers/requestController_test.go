package controllers_test

import (
	"berbagi/config"
	"berbagi/controllers"
	models "berbagi/models"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	config.InitDb()
	insertUser()
	insertAddress()
	insertDonor()
	insertChildren()
	insertFoundation()
	insertService()
	os.Exit(m.Run())
}

func insertUser() {
	user := []models.User{
		{
			ID:       1,
			Name:     "jono",
			RoleID:   2, // donor
			NIK:      "3333444455556666",
			Email:    "jono@jon.jon",
			Password: "jonjon",
		},
		{
			ID:       2,
			Name:     "jini",
			RoleID:   4, // children
			NIK:      "3333444455557777",
			Email:    "jini@jin.jin",
			Password: "jonjon",
		},
		{
			ID:       3,
			Name:     "joni",
			RoleID:   5, // foundation
			NIK:      "3333444455558888",
			Email:    "joni@jin.jin",
			Password: "jonjon",
		},
	}
	config.Db.Create(&user)
}

func insertAddress() {
	address := []models.Address{
		{
			ID:         1,
			Name:       "kebumen1",
			Latitude:   "-7.553644",
			Longitude:  "110,863470",
			CityID:     1,
			ProvinceID: 1,
		},
		{
			ID:         2,
			Name:       "kebumen2",
			Latitude:   "-7.553744",
			Longitude:  "110,863570",
			CityID:     1,
			ProvinceID: 1,
		},
		{
			ID:         3,
			Name:       "kebumen3",
			Latitude:   "-7.553844",
			Longitude:  "110,863670",
			CityID:     1,
			ProvinceID: 1,
		},
	}
	config.Db.Create(&address)
}

func insertDonor() {
	donor := []models.Donor{
		{
			UserID:    1,
			BirthDate: "1990-01-06",
			AddressID: 1,
		},
	}
	config.Db.Create(&donor)
}

func insertChildren() {
	children := []models.Children{
		{
			UserID:    2,
			BirthDate: "1990-01-06",
			AddressID: 2,
		},
	}
	config.Db.Create(&children)
}

func insertFoundation() {
	foundation := []models.Foundation{
		{
			UserID:    3,
			LicenseID: 231,
			AddressID: 3,
		},
	}
	config.Db.Create(&foundation)
}

func insertService() {
	foundation := []models.Proficiency{
		{
			Name: "health",
		},
	}
	config.Db.Create(&foundation)
}

func TestRequestGift(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		addressId            string
		packageId            string
		quantity             string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/request/gift",
			userId:               "2",
			role:                 "children",
			addressId:            "2",
			packageId:            "1",
			quantity:             "2",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "submitted",
			expectBodyContains2:  "School",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"address_id": testCase.addressId,
			"package_id": testCase.packageId,
			"quantity":   testCase.quantity,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.RequestGift(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestRequestDonation(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		foundationId         string
		role                 string
		addressId            string
		amount               string
		purpose              string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/request/donation",
			foundationId:         "3",
			role:                 "foundation",
			addressId:            "3",
			amount:               "100000",
			purpose:              "untuk",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "submitted",
			expectBodyContains2:  "100000",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		request := map[string]string{
			"address_id": testCase.addressId,
			"amount":     testCase.amount,
			"purpose":    testCase.purpose,
		}
		data, _ := json.Marshal(request)
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(string(data)))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Add("userId", testCase.foundationId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)

		if assert.NoError(t, controllers.RequestDonation(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestGetAllRequestList(t *testing.T) {
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
			path:                 "/request",
			userId:               "2",
			role:                 "children",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "gift",
		},
		{
			testName:             "success",
			path:                 "/request",
			userId:               "3",
			role:                 "foundation",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "service",
		},
		{
			testName:             "success",
			path:                 "/request?resolve=",
			userId:               "3",
			role:                 "foundation",
			resolved:             "yes",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "[]",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path + testCase.resolved)

		if assert.NoError(t, controllers.GetAllRequestListController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(testCase.testName, "all request ->", body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestGetTypeRequestList(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		role                 string
		addressId            string
		reqType              string
		resolved             string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/request/",
			userId:               "3",
			role:                 "foundation",
			addressId:            "3",
			reqType:              "service",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "start_date",
		},
		{
			testName:             "success",
			path:                 "/request/",
			userId:               "3",
			role:                 "foundation",
			addressId:            "3",
			reqType:              "donation",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "amount",
		},
		{
			testName:             "success",
			path:                 "/request/",
			userId:               "2",
			role:                 "children",
			addressId:            "2",
			reqType:              "gift",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "package",
		},
		{
			testName:             "success",
			path:                 "/request/",
			userId:               "2",
			role:                 "children",
			addressId:            "2",
			reqType:              "gift",
			resolved:             "?resolved=yes",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "[]",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		req.Header.Add("userId", testCase.userId)
		req.Header.Add("role", testCase.role)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path + testCase.reqType + testCase.resolved)

		if assert.NoError(t, controllers.GetTypeRequestListController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(testCase.testName, "type request ->", body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestGetRequestByRecipientId(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
		expectBodyContains2  string
	}{
		{
			testName:             "success",
			path:                 "/request/",
			userId:               "3",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "service",
		},
		{
			testName:             "success",
			path:                 "/request/",
			userId:               "2",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "list",
			expectBodyContains2:  "gift",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/", nil)
		req.Header.Add("userId", testCase.userId)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path + testCase.userId)

		if assert.NoError(t, controllers.GetRequestByRecipientIdController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			fmt.Println(testCase.testName, "user_id request ->", body)
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains2))
		}
	}
}

func TestDeleteRequest(t *testing.T) {
	var testCases = []struct {
		testName             string
		path                 string
		userId               string
		requestId            string
		expectStatus         int
		expectBodyStartsWith string
		expectBodyContains1  string
	}{
		{
			testName:             "success",
			path:                 "/request",
			userId:         "2",
			requestId: "1",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "delete",
		},
		{
			testName:             "failed",
			path:                 "/request",
			userId:         "2",
			requestId: "2",
			expectStatus:         http.StatusBadRequest,
			expectBodyStartsWith: "{\"status\":\"failed",
			expectBodyContains1:  "other's",
		},
		{
			testName:             "success",
			path:                 "/request",
			userId:         "3",
			requestId: "2",
			expectStatus:         http.StatusOK,
			expectBodyStartsWith: "{\"status\":\"success",
			expectBodyContains1:  "delete",
		},
	}

	e := echo.New()

	for _, testCase := range testCases {
		req := httptest.NewRequest(http.MethodDelete, "/", nil)
		req.Header.Set("userId", testCase.userId)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath(testCase.path)
		c.QueryParams().Add("request_id", testCase.requestId)

		if assert.NoError(t, controllers.DeleteRequestController(c)) {
			assert.Equal(t, testCase.expectStatus, rec.Code)
			body := rec.Body.String()
			assert.True(t, strings.HasPrefix(body, testCase.expectBodyStartsWith))
			assert.True(t, strings.Contains(body, testCase.expectBodyContains1))
		}
	}
}
