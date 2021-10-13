package controllers

import (
	"berbagi/config"
	"berbagi/models"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitLoginTest() *echo.Echo {
	config.InitDBTest("provinces", "cities", "addresses", "users", "admins", "donors", "volunteers", "childrens", "foundations", "roles")
	config.InsertProvince()
	config.InsertCity()
	config.InsertRole()
	e := echo.New()
	return e
}

func Test_LoginUserController(t *testing.T) {
	e := InitLoginTest()

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

	config.InsertUser(donorReqOK)

	marshalledDonorOK, _ := json.Marshal(donorReqOK)

	userReqInvalidEmail := models.User{
		Name:     "fattah",
		Email:    "",
		Password: "1234"}

	marshalledUserInvalidEmail, _ := json.Marshal(userReqInvalidEmail)

	userReqInvalidPassword := models.User{
		Name:     "abdul",
		Email:    "abdul@gmail.com",
		Password: "123"}

	marshalledUserInvalidPassword, _ := json.Marshal(userReqInvalidPassword)

	cases := []models.UserCaseWithBody{
		{
			Name:         "Valid login",
			Method:       "POST",
			Path:         "/login",
			ExpectedCode: http.StatusOK,
			RequestBody:  string(marshalledDonorOK),
			Message:      "You are logged in!"},
		{
			Name:         "Invalid login with invalid email",
			Method:       "POST",
			Path:         "/login",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  string(marshalledUserInvalidEmail),
			Message:      "No user with corresponding email"},
		{
			Name:         "Invalid login with invalid password",
			Method:       "POST",
			Path:         "/login",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  string(marshalledUserInvalidPassword),
			Message:      "Given password is incorrect"}}

	for _, testcase := range cases {
		req := httptest.NewRequest("POST", "/", strings.NewReader(testcase.RequestBody))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, LoginUserController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)
		}
	}
}
