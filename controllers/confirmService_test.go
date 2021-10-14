package controllers

import (
	"berbagi/config"
	"berbagi/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type ConfirmationResp struct {
	Status  string
	Message string
	Data    []models.ConfirmServicesAPI
}

func InitEchoTestConfirm() *echo.Echo {
	config.InitDbTest()
	e := echo.New()
	return e
}

func InsertServiceConfirm() error {
	confirm := []models.ConfirmServicesAPI{
		{
			VolunteerID: 1,
			UserID:      2,
			Invoice:     "001/BERBAGI/VOLUNTEER/001/003",
			StartDate:   FormatDate("2021-12-12"),
			FinishDate:  FormatDate("2021-12-30"),
		},
		{
			VolunteerID: 4,
			UserID:      3,
			Invoice:     "002/BERBAGI/VOLUNTEER/001/003",
			StartDate:   FormatDate("2021-12-12"),
			FinishDate:  FormatDate("2021-12-30"),
		},
	}

	tx := config.Db.Create(&confirm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func TestConfirmation(t *testing.T) {
	t.Run("get confirmation services", TestGetConfirmSevice)
	t.Run("Display verification letter", TestDisplayConfirmSevice)
	t.Run("add confirmation services", TestAddConfirmSevice)
}

func TestGetConfirmSevice(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
		id           string
	}{
		{
			name:         "unauthorized access",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "1",
		},
		{
			name:         "Success to get confirmation service data",
			expectedCode: http.StatusOK,
			sizeData:     1,
			id:           "1",
		},
		{
			name:         "Invalid verification id",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "s",
		},
		{
			name:         "verification id not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
			id:           "1",
		},
		{
			name:         "Failed to get confirmation service data",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "1",
		},
	}

	e := InitEchoTestConfirm()

	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()
	InsertDataUser()
	InsertDataVolunteer()
	InsertDataFoundation()

	for i := range testCases {
		if i == 3 {
			continue
		}

		if i == 1 {
			InsertServiceConfirm()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")
		userId2, _ := AddUser("annisa", "annisa@gmail.com", "annisa")

		req := httptest.NewRequest(http.MethodGet, "/services/verification/:verificationId", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if i == 3 {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId2)))
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		}

		c.SetParamNames("verificationId")
		c.SetParamValues(testCases[i].id)

		if i == 4 {
			config.Db.Migrator().DropTable(&models.ConfirmServicesAPI{})
		}

		if assert.NoError(t, GetConfirmServiceController(c)) {
			assert.Equal(t, testCases[i].expectedCode, rec.Code)
			body := rec.Body.String()

			var response VolunteerResp
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			switch i {
			case 0:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorize access", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get confirmation service data", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid verification id", response.Message)
			case 3:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "verification id not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get confirmation service data", response.Message)
			}
		}
	}
}

func TestDisplayConfirmSevice(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
		id           string
	}{
		{
			name:         "unauthorized access",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "1",
		},
		{
			name:         "Success to get confirmation service data",
			expectedCode: http.StatusOK,
			sizeData:     1,
			id:           "1",
		},
		{
			name:         "Invalid verification id",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "s",
		},
		{
			name:         "verification id not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
			id:           "1",
		},
		{
			name:         "Failed to get confirmation service data",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "1",
		},
	}

	e := InitEchoTestConfirm()

	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()
	InsertDataUser()
	InsertDataVolunteer()
	InsertDataFoundation()

	for i := range testCases {
		if i == 3 || i == 1 {
			continue
		}

		if i == 1 {
			InsertServiceConfirm()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")
		userId2, _ := AddUser("annisa", "annisa@gmail.com", "annisa")

		req := httptest.NewRequest(http.MethodGet, "/services/verification/:verificationId", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if i == 3 {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId2)))
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		}

		c.SetParamNames("verificationId")
		c.SetParamValues(testCases[i].id)

		if i == 4 {
			config.Db.Migrator().DropTable(&models.ConfirmServicesAPI{})
		}

		if assert.NoError(t, DisplayConfirmServiceController(c)) {
			assert.Equal(t, testCases[i].expectedCode, rec.Code)
			body := rec.Body.String()

			var response VolunteerResp
			err := json.Unmarshal([]byte(body), &response)
			if err != nil {
				assert.Error(t, err, "error")
			}

			switch i {
			case 0:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorize access", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get confirmation service data", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid verification id", response.Message)
			case 3:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "verification id not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get confirmation service data", response.Message)
			}
		}
	}
}

func TestAddConfirmSevice(t *testing.T) {
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
			name:         "Success to confirm service activity data",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
		{
			name:         "Failed to confim service activity data",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
	}

	e := InitEchoTestConfirm()

	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()
	InsertDataUser()
	InsertDataVolunteer()
	InsertDataFoundation()

	for i := range testCases {
		if i == 0 {
			continue
		}
		if i == 1 {
			InsertServiceCart()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")

		req := httptest.NewRequest(http.MethodPost, "/services/verification", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if i == 0 {
			c.Request().Header.Set("userId", "3")
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		}

		if i == 2 {
			config.Db.Migrator().DropTable(&models.ConfirmServicesAPI{})
		}

		if assert.NoError(t, AddConfirmServiceController(c)) {
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
				assert.Equal(t, "success to confirm service activity data", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to confim service activity data", response.Message)
			}
		}
	}
}
