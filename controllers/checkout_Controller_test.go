package controllers

import (
	"berbagi/config"
	"berbagi/models"

	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitCheckoutTest() *echo.Echo {
	config.InitDBTest("provinces", "cities", "addresses", "users", "donors", "childrens", "payment_methods",
		"categories", "products", "product_packages", "product_package_details", "product_carts", "transactions", "transaction_details")

	config.InsertProvince()
	config.InsertCity()
	config.InsertCategory()
	config.InsertProduct()
	config.InsertProductPackage()
	config.InsertProductPackageDetail()
	config.InsertPaymentMethod()

	e := echo.New()
	return e
}

func Test_GetCheckoutByUserIdController(t *testing.T) {
	e := InitCheckoutTest()

	donor := models.RegistrationAPI{
		Name:        "abdul",
		Email:       "abdul@gmail.com",
		Password:    "1234",
		NIK:         "123450",
		AddressName: "Rumah Abdul",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "123,111",
		Latitude:    "111,76",
		RoleID:      2}

	children1 := models.RegistrationAPI{
		Name:        "dara",
		Email:       "dara@gmail.com",
		Password:    "1234",
		NIK:         "123451",
		AddressName: "Rumah Dara",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "125,111",
		Latitude:    "114,76",
		RoleID:      4}

	children2 := models.RegistrationAPI{
		Name:        "ali",
		Email:       "ali@gmail.com",
		Password:    "1234",
		NIK:         "123452",
		AddressName: "Rumah Ali",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "128,111",
		Latitude:    "118,76",
		RoleID:      4}

	donorID, _ := config.InsertUser(donor)
	children1ID, _ := config.InsertUser(children1)
	children2ID, _ := config.InsertUser(children2)

	emptyCarts := models.UserCaseWithBody{
		Name:         "Get checkout with empty cart",
		Method:       "GET",
		Path:         "/checkout",
		ExpectedCode: http.StatusBadRequest,
		RequestBody:  "",
		Message:      "no product_package_id found in user's product_carts",
		Size:         0}

	reqOut := httptest.NewRequest(emptyCarts.Method, "/", nil)
	reqOut.Header.Set("Content-Type", "application/json")
	reqOut.Header.Set("userId", fmt.Sprintf("%v", donorID))
	recOut := httptest.NewRecorder()
	cOut := e.NewContext(reqOut, recOut)

	cOut.SetPath(emptyCarts.Path)

	if assert.NoError(t, GetCheckoutByUserIdController(cOut)) {
		assert.Equal(t, emptyCarts.ExpectedCode, recOut.Code)

		var userResponseOut models.ResponseOK

		if err := json.Unmarshal([]byte(recOut.Body.String()), &userResponseOut); err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, emptyCarts.Message, userResponseOut.Message)

		userData, _ := userResponseOut.Data.(map[string]interface{})
		recipients, _ := userData["recipients"].([]interface{})
		assert.Equal(t, emptyCarts.Size, len(recipients))
	}

	cart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 1, Quantity: 2},
		{RecipientID: uint(children1ID), ProductPackageID: 3, Quantity: 5}, {RecipientID: uint(children2ID), ProductPackageID: 2, Quantity: 1}}

	config.InsertProductCart(cart, donorID)

	cases := []models.UserCaseWithBody{
		{
			Name:         "Get checkout from non empty product cart",
			Method:       "POST",
			Path:         "/checkout",
			ExpectedCode: http.StatusOK,
			RequestBody:  "",
			Message:      "cart is retrieved succesfully!",
			Size:         2}}

	for _, testcase := range cases {
		req := httptest.NewRequest(testcase.Method, "/", nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("userId", fmt.Sprintf("%v", donorID))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, GetCheckoutByUserIdController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)
			userData, _ := userResponse.Data.(map[string]interface{})
			recipients, _ := userData["recipients"].([]interface{})
			assert.Equal(t, testcase.Size, len(recipients))

		}
	}

}

func Test_AddCheckoutByUserIdController(t *testing.T) {
	e := InitCheckoutTest()

	donor := models.RegistrationAPI{
		Name:        "abdul",
		Email:       "abdul@gmail.com",
		Password:    "1234",
		NIK:         "123450",
		AddressName: "Rumah Abdul",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "123,111",
		Latitude:    "111,76",
		RoleID:      2}

	children1 := models.RegistrationAPI{
		Name:        "dara",
		Email:       "dara@gmail.com",
		Password:    "1234",
		NIK:         "123451",
		AddressName: "Rumah Dara",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "125,111",
		Latitude:    "114,76",
		RoleID:      4}

	children2 := models.RegistrationAPI{
		Name:        "ali",
		Email:       "ali@gmail.com",
		Password:    "1234",
		NIK:         "123452",
		AddressName: "Rumah Ali",
		ProvinceID:  1,
		CityID:      1,
		Longitude:   "128,111",
		Latitude:    "118,76",
		RoleID:      4}

	donorID, _ := config.InsertUser(donor)
	children1ID, _ := config.InsertUser(children1)
	children2ID, _ := config.InsertUser(children2)

	validPaymentMethod := models.PaymentMethod{ID: uint(1)}

	var validData bytes.Buffer
	json.NewEncoder(&validData).Encode(validPaymentMethod)

	invalidPaymentMethod := models.PaymentMethod{ID: uint(122)}

	var invalidData1 bytes.Buffer
	json.NewEncoder(&invalidData1).Encode(invalidPaymentMethod)

	emptyCarts := models.UserCaseWithBody{
		Name:         "Add checkout with empty cart",
		Method:       "POST",
		Path:         "/checkout",
		ExpectedCode: http.StatusBadRequest,
		RequestBody:  validData.String(),
		Message:      "no product_package found in the cart. add product_package first before checking out"}

	reqOut := httptest.NewRequest(emptyCarts.Method, "/", strings.NewReader(emptyCarts.RequestBody))
	reqOut.Header.Set("Content-Type", "application/json")
	reqOut.Header.Set("userId", fmt.Sprintf("%v", donorID))
	recOut := httptest.NewRecorder()
	cOut := e.NewContext(reqOut, recOut)

	cOut.SetPath(emptyCarts.Path)

	if assert.NoError(t, AddCheckoutByUserIdController(cOut)) {
		assert.Equal(t, emptyCarts.ExpectedCode, recOut.Code)

		var userResponseOut models.ResponseOK

		if err := json.Unmarshal([]byte(recOut.Body.String()), &userResponseOut); err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, emptyCarts.Message, userResponseOut.Message)

	}

	cart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 1, Quantity: 2},
		{RecipientID: uint(children1ID), ProductPackageID: 3, Quantity: 5}, {RecipientID: uint(children2ID), ProductPackageID: 2, Quantity: 1}}

	config.InsertProductCart(cart, donorID)

	cases := []models.UserCaseWithBody{
		{
			Name:         "Add checkout with invalid payment method",
			Method:       "POST",
			Path:         "/checkout",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  invalidData1.String(),
			Message:      "no payment_method_id '122' found in the payment method table"},
		{
			Name:         "Add checkout with valid payment method",
			Method:       "POST",
			Path:         "/checkout",
			ExpectedCode: http.StatusOK,
			RequestBody:  validData.String(),
			Message:      "checkout is succesfull"}}

	for _, testcase := range cases {
		req := httptest.NewRequest(testcase.Method, "/", strings.NewReader(testcase.RequestBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("userId", fmt.Sprintf("%v", donorID))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, AddCheckoutByUserIdController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)

		}
	}

}
