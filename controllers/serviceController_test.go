package controllers

import (
	"berbagi/config"
	"berbagi/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type ServiceResp struct {
	Status  string
	Message string
	Data    []models.Service
}

func InitEchoTestService() *echo.Echo {
	config.InitDbTest()
	e := echo.New()
	return e
}

func InsertDataFoundation() error {
	user := []models.Foundation{
		{
			UserID:    2,
			LicenseID: 123,
			AddressID: 2,
		},
		{
			UserID:    3,
			LicenseID: 1234,
			AddressID: 3,
		},
	}

	tx := config.Db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func InsertServiceCart() error {
	service := models.ServiceCart{
		VolunteerID: 1,
		UserID:      2,
		AddressID:   2,
	}

	startDate := FormatDate("2021-12-12")
	finishDate := FormatDate("2021-12-30")

	service.StartDate = startDate
	service.FinishDate = finishDate

	tx := config.Db.Create(&service)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func FormatDate(date string) time.Time {
	formatedDate, _ := time.Parse("2006-01-02", date)
	return formatedDate
}

func TestServices(t *testing.T) {
	t.Run("get service on cart", TestGetServiceOnCart)
	t.Run("add service to cart", TestAddServiceToCart)
	t.Run("update service on cart", TestUpdateServiceOnCart)
	t.Run("delete service on cart", TestDeleteServiceOnCart)
}

func TestGetServiceOnCart(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "service cart not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
		},
		{
			name:         "success get service cart",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
		{
			name:         "failed to get service cart",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
	}

	e := InitEchoTestService()

	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()
	InsertDataUser()
	InsertDataVolunteer()
	InsertDataFoundation()

	for i := range testCases {
		if i == 1 {
			InsertServiceCart()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")

		req := httptest.NewRequest(http.MethodGet, "/services", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i == 2 {
			config.Db.Migrator().DropTable(&models.ServiceCart{})
		}

		if assert.NoError(t, GetServiceOnCartController(c)) {
			assert.Equal(t, testCases[i].expectedCode, rec.Code)
			body := rec.Body.String()

			var response VolunteerResp
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			switch i {
			case 0:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "service cart not found !", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get service cart", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get service cart", response.Message)
			}
		}
	}
}

func TestAddServiceToCart(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		id           string
	}{
		{
			name:         "volunteer id not found",
			expectedCode: http.StatusOK,
			id:           "5",
		},
		{
			name:         "success to add service on cart",
			expectedCode: http.StatusOK,
			id:           "1",
		},
		{
			name:         "find another date!",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
		{
			name:         "failed to add service to cart",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
	}

	e := InitEchoTestProficiency()

	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()
	InsertDataUser()
	InsertDataVolunteer()
	InsertDataFoundation()

	for i := range testCases {

		var startDate string

		if i == 1 {
			startDate = "2021-12-13"
		} else {
			startDate = "2021-10-01"
		}

		reqBody := map[string]interface{}{
			"start_date":   startDate,
			"finish_date":  "2022-01-12",
			"recipient_id": 3,
			"address_id":   3,
		}

		requestBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/services", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")

		if i == 0 {
			c.Request().Header.Set("userId", testCases[0].id)
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		}

		if i == 3 {
			config.Db.Migrator().DropTable(&models.Foundation{})
		}

		c.SetParamNames("id")
		c.SetParamValues(testCases[i].id)

		if assert.NoError(t, AddServiceToCartController(c)) {

			assert.Equal(t, testCases[i].expectedCode, rec.Code)

			body := rec.Body.String()

			var response ProficiencyResp
			err := json.Unmarshal([]byte(body), &response)

			if err != nil {
				assert.Error(t, err, "error")
			}

			switch i {
			case 0:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "volunteer id not found", response.Message)
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to add service on cart", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "find another date!", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to add service to cart", response.Message)
			}
		}
	}
}

func TestDeleteServiceOnCart(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "volunteer id not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
		},
		{
			name:         "success to delete service on cart",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
		{
			name:         "failed to delete service cart",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
	}

	e := InitEchoTestService()

	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()
	InsertDataUser()
	InsertDataVolunteer()
	InsertDataFoundation()

	for i := range testCases {
		if i == 1 {
			InsertServiceCart()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")

		req := httptest.NewRequest(http.MethodDelete, "/services", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i == 2 {
			config.Db.Migrator().DropTable(&models.ServiceCart{})
		}

		if assert.NoError(t, DeleteServiceCartController(c)) {
			assert.Equal(t, testCases[i].expectedCode, rec.Code)
			body := rec.Body.String()

			var response VolunteerResp
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			switch i {
			case 0:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "volunteer id not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to delete service on cart", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to delete service cart", response.Message)
			}
		}
	}
}

func TestUpdateServiceOnCart(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		id           string
	}{
		{
			name:         "volunteer id not found",
			expectedCode: http.StatusOK,
			id:           "3",
		},
		{
			name:         "success to update service on carts",
			expectedCode: http.StatusOK,
			id:           "1",
		},
		{
			name:         "find another date!",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
		{
			name:         "failed to update service cart",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
	}

	e := InitEchoTestProficiency()
	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()
	InsertDataUser()
	InsertDataVolunteer()
	InsertDataFoundation()

	for i := range testCases {

		if i == 1 {
			InsertServiceCart()
		}

		var startDate string

		if i == 1 {
			startDate = "2021-12-13"
		} else {
			startDate = "2021-10-01"
		}

		reqBody := map[string]interface{}{
			"start_date":   startDate,
			"finish_date":  "2022-01-12",
			"recipient_id": 3,
			"address_id":   3,
		}

		requestBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPut, "/services", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")
		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i == 3 {
			config.Db.Migrator().DropTable(&models.ServiceCart{})
		}

		c.SetParamNames("id")
		c.SetParamValues(testCases[i].id)

		if assert.NoError(t, UpdatedServiceOncartController(c)) {

			assert.Equal(t, testCases[i].expectedCode, rec.Code)

			body := rec.Body.String()

			var response ProficiencyResp
			err := json.Unmarshal([]byte(body), &response)

			if err != nil {
				assert.Error(t, err, "error")
			}

			switch i {
			case 0:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "volunteer id not found", response.Message)
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to update service on carts", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "find another date!", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to update service cart", response.Message)
			}
		}
	}
}
