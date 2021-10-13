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

type CertificateResp struct {
	Status  string
	Message string
	Data    models.Certificate
}

func InitEchoTestCertificate() *echo.Echo {
	config.InitDbTest()
	e := echo.New()
	return e
}

func InsertCertificate() error {
	certificate := []models.Certificate{
		{
			CompletionID: 1,
		},
		{
			CompletionID: 2,
		},
	}

	tx := config.Db.Create(&certificate)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func TestCertificate(t *testing.T) {
	t.Run("get certificate of completion", TestGetCertificate)
	t.Run("Display certificate", TestDisplayCertificate)
}

func TestGetCertificate(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
		id           string
	}{
		{
			name:         "completion id not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
			id:           "2",
		},
		{
			name:         "success to get certificate",
			expectedCode: http.StatusOK,
			sizeData:     1,
			id:           "1",
		},
		{
			name:         "Invalid completion id",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "s",
		},
		{
			name:         "unauthorized access",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "1",
		},
		{
			name:         "Failed to get completion data",
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
	InsertServiceCart()
	InsertServiceConfirm()

	for i := range testCases {
		if i == 3 {
			continue
		}
		if i == 1 {
			InsertCompletion()
			InsertCertificate()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")
		userId2, _ := AddUser("kiki", "kiki@gmail.com", "kiki")

		req := httptest.NewRequest(http.MethodGet, "/certificates/:completionId", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if i == 0 {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId2)))
		} else if i == 3 {
			c.Request().Header.Add("userId", "2")
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		}

		c.SetParamNames("completionId")
		c.SetParamValues(testCases[i].id)

		if i == 4 {
			config.Db.Migrator().DropTable(&models.ConfirmServicesAPI{})
		}

		if assert.NoError(t, GetCertificateController(c)) {
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
				assert.Equal(t, "completion id not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get certificate", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid completion id", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get certificate", response.Message)
			}
		}
	}
}

func TestDisplayCertificate(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
		id           string
	}{
		{
			name:         "completion id not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
			id:           "2",
		},
		{
			name:         "success to get certificate",
			expectedCode: http.StatusOK,
			sizeData:     1,
			id:           "1",
		},
		{
			name:         "Invalid completion id",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "s",
		},
		{
			name:         "unauthorized access",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
			id:           "1",
		},
		{
			name:         "Failed to get completion data",
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
	InsertServiceCart()
	InsertServiceConfirm()

	for i := range testCases {
		if i == 3 {
			continue
		}
		if i == 2 {
			InsertCompletion()
			InsertCertificate()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")
		userId2, _ := AddUser("kiki", "kiki@gmail.com", "kiki")

		req := httptest.NewRequest(http.MethodGet, "/certificates/display/:completionId", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if i == 0 {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId2)))
		} else if i == 3 {
			c.Request().Header.Add("userId", "2")
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		}

		c.SetParamNames("completionId")
		c.SetParamValues(testCases[i].id)

		if i == 4 {
			config.Db.Migrator().DropTable(&models.ConfirmServicesAPI{})
		}

		if assert.NoError(t, CertificateDisplayController(c)) {
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
				assert.Equal(t, "completion id not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				// assert.Equal(t, "success to get certificate", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid completion id", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get certificate", response.Message)
			}
		}
	}
}
