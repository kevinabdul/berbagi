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

type VolunteerResp struct {
	Status  string
	Message string
	Data    []models.User
}

func InitEchoTestVolunteer() *echo.Echo {
	config.InitDbTest()
	e := echo.New()
	return e
}

func InsertDataUser() error {
	user := []models.User{
		{
			Name:     "rokhiyah",
			NIK:      "3507652897537",
			Email:    "rokhiyah@gmail.com",
			Password: "rokhiyah",
			RoleID:   3,
		},
		{
			Name:     "yayasan",
			NIK:      "3507652897539",
			Email:    "yayasan@gmail.com",
			Password: "yayasan",
			RoleID:   5,
		},
		{
			Name:     "yayasan 1",
			NIK:      "3507652897531",
			Email:    "yayasan1@gmail.com",
			Password: "yayasan1",
			RoleID:   5,
		},
		{
			Name:     "annisa",
			NIK:      "3507652897532",
			Email:    "annisa@gmail.com",
			Password: "annisa",
			RoleID:   3,
		},
		{
			Name:     "kiki",
			NIK:      "3507152897532",
			Email:    "kiki@gmail.com",
			Password: "kiki",
			RoleID:   3,
		},
	}
	tx := config.Db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func InsertDataVolunteer() error {
	user := []models.Volunteer{
		{
			UserID:        1,
			ProficiencyID: 1,
			AddressID:     1,
		},
		{
			UserID:        4,
			ProficiencyID: 1,
			AddressID:     1,
		},
	}

	tx := config.Db.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func InsertAddress() error {
	address := []models.Address{
		{
			Name:       "flamboyan",
			CityID:     1,
			ProvinceID: 1,
		},
		{
			Name:       "cempaka",
			CityID:     1,
			ProvinceID: 1,
		},
		{
			Name:       "kenanga",
			CityID:     1,
			ProvinceID: 1,
		},
	}

	tx := config.Db.Create(&address)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func InsertProvince() {
	provinces := []models.Province{{Name: "DKI Jakarta"}, {Name: "Denpasar"}}

	config.Db.Create(&provinces)
}

func InsertCity() {
	cities := []models.City{{Name: "Jakarta Pusat", ProvinceID: 1}, {Name: "Bali", ProvinceID: 2}}

	config.Db.Create(&cities)
}

func TestVolunteer(t *testing.T) {
	t.Run("get list volunteer", TestGetListVolunteer)
	t.Run("get volunteer profile", TestGetVolunteerProfile)
}

func TestGetListVolunteer(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "list volunteer not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
		},
		{
			name:         "success get list of volunteers",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
		{
			name:         "unauthorized access to get list of volunteer",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
		{
			name:         "failed to get list of volunteer",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
	}

	e := InitEchoTestVolunteer()

	InsertDataProficiency()

	for i := range testCases {
		if i == 1 {
			InsertDataUser()
		}

		userId, _ := AddUser("urnik", "urnikrokhiyah@gmail.com", "12345")

		req := httptest.NewRequest(http.MethodGet, "/volunteers", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.Request().Header.Set("userId", strconv.Itoa(int(userId)))

		if i != 2 {
			c.Request().Header.Set("role", "admin")
		}

		if i == 3 {
			config.Db.Migrator().DropTable(&models.User{})
		}

		if assert.NoError(t, GetListVolunteersController(c)) {
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
				assert.Equal(t, "volunteer data not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get list volunteers", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "unauthorized access", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get list volunteers", response.Message)
			}
		}
	}
}

func TestGetVolunteerProfile(t *testing.T) {
	var testCases = []struct {
		name         string
		expectedCode int
		sizeData     int
	}{
		{
			name:         "volunteer not found",
			expectedCode: http.StatusOK,
			sizeData:     0,
		},
		{
			name:         "success get volunteer profile",
			expectedCode: http.StatusOK,
			sizeData:     1,
		},
		{
			name:         "invalid volunteer id",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
		{
			name:         "failed to get volunteer profile",
			expectedCode: http.StatusBadRequest,
			sizeData:     0,
		},
	}

	e := InitEchoTestVolunteer()

	InsertDataProficiency()
	InsertProvince()
	InsertCity()
	InsertAddress()

	for i := range testCases {
		if i == 1 {
			InsertDataUser()
			InsertDataVolunteer()
		}

		req := httptest.NewRequest(http.MethodGet, "/volunteers/profile", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		if i == 2 {
			c.Request().Header.Set("userId", "s")
		} else {
			c.Request().Header.Set("userId", strconv.Itoa(int(1)))
		}

		if i == 2 {
			config.Db.Migrator().DropTable(&models.User{})
		}

		if assert.NoError(t, GetVolunteerProfileController(c)) {
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
				assert.Equal(t, "volunteer data not found", response.Message)
				assert.Equal(t, testCases[i].sizeData, len(response.Data))
			case 1:
				assert.Equal(t, "success", response.Status)
				assert.Equal(t, "success to get volunteer profile", response.Message)
			case 2:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "invalid volunteer id", response.Message)
			case 3:
				assert.Equal(t, "failed", response.Status)
				assert.Equal(t, "failed to get volunteer profile", response.Message)
			}
		}
	}
}
