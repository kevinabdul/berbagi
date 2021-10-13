package controllers

import (
	"berbagi/config"
	"berbagi/models"
	"berbagi/utils/password"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

type ProficiencyResp struct {
	Status  string
	Message string
	Data    []models.Proficiency
}

func InitEchoTestProficiency() *echo.Echo {
	config.InitDbTest()
	e := echo.New()
	return e
}

func InsertDataProficiency() error {
	proficiency := models.Proficiency{
		Name: "education",
	}

	tx := config.Db.Create(&proficiency)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func AddUser(name, email, userPassword string) (uint, error) {
	pass, _ := password.Hash(userPassword)
	user := models.User{Name: name, Email: email, Password: pass}
	res := config.Db.FirstOrCreate(&user)

	if res.Error != nil {
		return uint(0), res.Error
	}
	return user.ID, nil
}

func TestProficiency(t *testing.T) {
	t.Run("get list proficiency", TestGetAllProficienciesController)
	t.Run("create new proficiency", TestCreateNewProficiencyController)
	t.Run("delete proficiency", TestDeleteProficiency)
	t.Run("update proficiency data", TestUpdateProficiency)
}

func TestCreateNewProficiencyController(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "success create new proficiency data",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
		{
			name:         "unauthorized access to create new proficiency data",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
		{
			name:         "failed to create proficiency data",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
	}

	e := InitEchoTestProficiency()

	for i := range testCases {

		reqBody := map[string]string{
			"name": "education",
		}

		requestBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPost, "/proficiencies", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		userId, _ := AddUser("urnik", "urnikrokhiyah@gmail.com", "12345")
		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i != 1 {
			c.Request().Header.Set("role", "admin")
		}

		if i == 2 {
			config.Db.Migrator().DropTable(&models.Proficiency{})
		}

		if assert.NoError(t, CreateNewProficiencyController(c)) {
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
				assert.Equal(t, "success to create new proficiency", response.Message)
			case 1:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to create new proficiency", response.Message)
			}
		}
	}
}

func TestGetAllProficienciesController(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "list proficiencies not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
		},
		{
			name:         "success get proficiency data",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
		{
			name:         "unauthorized access to get list proficiency data",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
		{
			name:         "failed to get proficiency data",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
	}

	e := InitEchoTestProficiency()

	for i := range testCases {
		if i == 1 {
			InsertDataProficiency()
		}

		userId, _ := AddUser("urnik", "urnikrokhiyah@gmail.com", "12345")

		req := httptest.NewRequest(http.MethodGet, "/proficiencies", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i != 2 {
			c.Request().Header.Set("role", "admin")
		}

		if i == 3 {
			config.Db.Migrator().DropTable(&models.Proficiency{})
		}

		if assert.NoError(t, GetAllProficienciesController(c)) {
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
				assert.Equal(t, "list proficiencies not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get list proficiencies", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get all list proficiencies", response.Message)
			}
		}
	}
}

func TestDeleteProficiency(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		id           string
	}{
		{
			name:         "proficiency data not found",
			expectedCode: http.StatusOK,
			id:           "2",
		},
		{
			name:         "success delete proficiency data",
			expectedCode: http.StatusOK,
			id:           "1",
		},
		{
			name:         "unauthorized access",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
		{
			name:         "invalid proficiency id",
			expectedCode: http.StatusBadRequest,
			id:           "s",
		},
		{
			name:         "failed to delete proficiency data",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
	}

	e := InitEchoTestProficiency()

	for i := range testCases {

		InsertDataProficiency()

		req := httptest.NewRequest(http.MethodDelete, "/proficiencies", nil)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		userId, _ := AddUser("urnik", "urnikrokhiyah@gmail.com", "12345")
		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i != 2 {
			c.Request().Header.Set("role", "admin")
		}

		if i == 4 {
			config.Db.Migrator().DropTable(&models.Proficiency{})
		}

		c.SetParamNames("id")
		c.SetParamValues(testCases[i].id)

		if assert.NoError(t, DeleteProficiencyController(c)) {
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
				assert.Equal(t, "proficiency not found", response.Message)
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to delete proficiency", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid proficiency id", response.Message)
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to delete proficiency", response.Message)
			}
		}
	}
}

func TestUpdateProficiency(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		id           string
	}{
		{
			name:         "proficiency data not found",
			expectedCode: http.StatusOK,
			id:           "2",
		},
		{
			name:         "success update proficiency data",
			expectedCode: http.StatusOK,
			id:           "1",
		},
		{
			name:         "unauthorized access",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
		{
			name:         "invalid proficiency id",
			expectedCode: http.StatusBadRequest,
			id:           "s",
		},
		{
			name:         "failed to update proficiency data",
			expectedCode: http.StatusBadRequest,
			id:           "1",
		},
	}

	e := InitEchoTestProficiency()

	for i := range testCases {

		if i != 0 {
			InsertDataProficiency()
		}

		reqBody := map[string]string{
			"name": "education",
		}

		requestBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest(http.MethodPut, "/proficiencies", bytes.NewBuffer(requestBody))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		c := e.NewContext(req, rec)

		userId, _ := AddUser("urnik", "urnikrokhiyah@gmail.com", "12345")
		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i != 2 {
			c.Request().Header.Set("role", "admin")
		}

		if i == 4 {
			config.Db.Migrator().DropTable(&models.Proficiency{})
		}

		c.SetParamNames("id")
		c.SetParamValues(testCases[i].id)

		if assert.NoError(t, UpdatedProficiencyController(c)) {
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
				assert.Equal(t, "proficiency not found", response.Message)
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to update proficiency", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid proficiency id", response.Message)
			case 4:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to update proficiency", response.Message)
			}
		}
	}
}
