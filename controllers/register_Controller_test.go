package controllers

import (
	"berbagi/config"
	"berbagi/models"

	"net/http"
	"testing"

	//"net/url"
	"encoding/json"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitRegisterTest() *echo.Echo {
	config.InitDBTest("provinces", "cities", "addresses", "users", "admins", "donors", "volunteers", "childrens", "foundations")
	config.InsertProvince()
	config.InsertCity()
	e := echo.New()
	return e
}

func Test_RegisterUserController(t *testing.T) {
	e := InitRegisterTest()

	donorReqOK := models.RegistrationAPI{
		Name:        "abdul",
		Email:       "abdul@gmail.com",
		Password:    "1234",
		NIK:         "12345",
		AddressName: "Rumah Abdul",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "123,111",
		Latitude:    "111,76",
		RoleID:      2}

	marshalledDonorOk, _ := json.Marshal(donorReqOK)

	donorReqInvalidEmail := models.RegistrationAPI{
		Name:        "abdul",
		Email:       "",
		Password:    "1234",
		NIK:         "123456",
		AddressName: "Rumah Abdul",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "123,111",
		Latitude:    "111,76",
		RoleID:      2}

	marshalledDonorInvalidEmail, _ := json.Marshal(donorReqInvalidEmail)

	donorReqInvalidNIK := models.RegistrationAPI{
		Name:        "abdul razi",
		Email:       "abdul.razi@gmail.com",
		Password:    "1234",
		NIK:         "12345",
		AddressName: "Rumah Abdul",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "123,111",
		Latitude:    "111,76",
		RoleID:      2}

	marshalledDonorInvalidNIK, _ := json.Marshal(donorReqInvalidNIK)

	cases := []models.UserCaseWithBody{
		{
			Name:         "Add user",
			Method:       "POST",
			Path:         "/register",
			ExpectedCode: http.StatusOK,
			RequestBody:  string(marshalledDonorOk),
			Message:      "user has been created!"},
		{
			Name:         "Add user with invalid email",
			Method:       "POST",
			Path:         "/register",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  string(marshalledDonorInvalidEmail),
			Message:      "Invalid Email or Password. Make sure its not empty and are of string type"},
		{
			Name:         "Add user with same NIK",
			Method:       "POST",
			Path:         "/register",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  string(marshalledDonorInvalidNIK),
			Message:      "Error 1062: Duplicate entry '12345' for key 'nik'"}}

	for _, testcase := range cases {
		req := httptest.NewRequest("POST", "/", strings.NewReader(testcase.RequestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, RegisterUserController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)
		}
	}
}
