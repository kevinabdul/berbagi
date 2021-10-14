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

type CompletionResp struct {
	Status  string
	Message string
	Data    models.Completion
}

func InitEchoTestCompletion() *echo.Echo {
	config.InitDbTest()
	e := echo.New()
	return e
}

func InsertCompletion() error {
	completion := []models.Completion{
		{
			ConfirmServicesAPIID: 1,
			CompletionStatus:     "completed",
		},
		{
			ConfirmServicesAPIID: 2,
			CompletionStatus:     "completed",
		},
	}

	tx := config.Db.Create(&completion)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func TestCompletion(t *testing.T) {
	t.Run("get completion", TestGetCompletion)
	t.Run("update completion status", TestUpdateCompletionStatus)
}

func TestGetCompletion(t *testing.T) {
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
			name:         "Success to get completion data",
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
			id:           "3",
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
		if i != 3 {
			InsertCompletion()
		}

		userId, _ := AddUser("rokhiyah", "rokhiyah@gmail.com", "rokhiyah")
		userId2, _ := AddUser("annisa", "annisa@gmail.com", "annisa")

		req := httptest.NewRequest(http.MethodGet, "/completion/:verificationId", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if i == 3 {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId2)))
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(userId)))
		}

		if i != 0 {
			c.Request().Header.Set("role", "admin")
		}

		c.SetParamNames("verificationId")
		c.SetParamValues(testCases[i].id)

		if i == 4 {
			config.Db.Migrator().DropTable(&models.ConfirmServicesAPI{})
		}

		if assert.NoError(t, GetCompletionDetailController(c)) {
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
				assert.Equal(t, "unauthorized access", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get completion data", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid verification id", response.Message)
			case 3:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "verification id not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get completion data", response.Message)
			}
		}
	}
}

func TestUpdateCompletionStatus(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		id           string
	}{
		{
			name:         "verification id not found",
			expectedCode: http.StatusOK,
			id:           "2",
		},
		{
			name:         "success to update completion status",
			expectedCode: http.StatusOK,
			id:           "1",
		},
		{
			name:         "Unauthorized access",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
		{
			name:         "Invalid verification id",
			expectedCode: http.StatusBadRequest,
			id:           "s",
		},
		{
			name:         "Invalid completion status",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
		{
			name:         "Failed to update completion status",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
	}

	e := InitEchoTestCompletion()
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

		if i != 0 {
			InsertCompletion()
		}

		req := httptest.NewRequest(http.MethodPut, "/completion/:verificationId", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		userId, _ := AddUser("urnik", "urnikrokhiyah@gmail.com", "12345")
		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i != 2 {
			c.Request().Header.Set("role", "admin")
		}

		if i == 5 {
			config.Db.Migrator().DropTable(&models.Completion{})
		}

		c.SetParamNames("verificationId")
		c.SetParamValues(testCases[i].id)

		if i == 4 {
			c.QueryParams().Add("status", "finished")
		} else {
			c.QueryParams().Add("status", "verified")
		}

		if assert.NoError(t, UpdateStatusCompletionController(c)) {
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
				assert.Equal(t, "verification id not found", response.Message)
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to update completion status", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid verification id", response.Message)
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid completion status", response.Message)
			case 5:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to update completion status", response.Message)
			}
		}
	}
}
