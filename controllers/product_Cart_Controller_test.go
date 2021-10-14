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

func InitProductCartTest() *echo.Echo {
	config.InitDBTest("provinces", "cities", "addresses", "users", "donors", "childrens",
		"categories", "products", "product_packages", "product_package_details", "product_carts")

	config.InsertProvince()
	config.InsertCity()
	config.InsertCategory()
	config.InsertProduct()
	config.InsertProductPackage()
	config.InsertProductPackageDetail()

	e := echo.New()
	return e
}

func Test_GetProductCartByUserIdController(t *testing.T) {
	e := InitProductCartTest()

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

	emptyCase := models.UserCaseWithBody{
		Name:         "Get items from user's product cart that's empty",
		Method:       "GET",
		Path:         "/product-carts",
		ExpectedCode: http.StatusBadRequest,
		RequestBody:  "",
		Message:      "no product_package_id found in user's product_carts"}

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(emptyCase.Path)

	if assert.NoError(t, GetProductCartByUserIdController(c)) {
		assert.Equal(t, emptyCase.ExpectedCode, rec.Code)

		var userResponse models.ResponseOK

		if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, emptyCase.Message, userResponse.Message)
	}

	userCart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 1, Quantity: 2},
		{RecipientID: uint(children1ID), ProductPackageID: 3, Quantity: 5}, {RecipientID: uint(children2ID), ProductPackageID: 2, Quantity: 1}}

	config.InsertProductCart(userCart, donorID)

	cases := []models.UserCaseWithBody{
		{
			Name:         "Get item from user's product carts",
			Method:       "GET",
			Path:         "/product-carts",
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

		if assert.NoError(t, GetProductCartByUserIdController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)

			data, _ := userResponse.Data.(map[string]interface{})

			packages, _ := data["recipients"].([]interface{})

			assert.Equal(t, testcase.Size, len(packages))

		}
	}
}

func Test_UpdateProductCartByUserIdController(t *testing.T) {
	e := InitProductCartTest()

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

	validCart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 1, Quantity: 2},
		{RecipientID: uint(children1ID), ProductPackageID: 3, Quantity: 5}, {RecipientID: uint(children2ID), ProductPackageID: 2, Quantity: 1}}

	var validData bytes.Buffer
	json.NewEncoder(&validData).Encode(validCart)

	invalidCart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 789, Quantity: 2}}

	var invalidData1 bytes.Buffer
	json.NewEncoder(&invalidData1).Encode(invalidCart)

	invalidRecipient := []models.ProductCart{{RecipientID: uint(234567), ProductPackageID: 1, Quantity: 2}}

	var invalidData2 bytes.Buffer
	json.NewEncoder(&invalidData2).Encode(invalidRecipient)

	selfDonate := []models.ProductCart{{RecipientID: uint(donorID), ProductPackageID: 1, Quantity: 2}}

	var invalidData3 bytes.Buffer
	json.NewEncoder(&invalidData3).Encode(selfDonate)

	cases := []models.UserCaseWithBody{
		{
			Name:         "Update Item with valid product package id",
			Method:       "PUT",
			Path:         "/product-carts",
			ExpectedCode: http.StatusOK,
			RequestBody:  validData.String(),
			Message:      "product cart is updated!"},
		{
			Name:         "Update Item with invalid product package id",
			Method:       "PUT",
			Path:         "/product-carts",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  invalidData1.String(),
			Message:      "no product_package_id with id: 789 found in the product_package table"},
		{
			Name:         "Update Item with recipient id that doesnt exist",
			Method:       "PUT",
			Path:         "/product-carts",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  invalidData2.String(),
			Message:      "no recipient_id with id: 234567 found in the children table"},
		{
			Name:         "Update Item using same donor id and recipient id",
			Method:       "PUT",
			Path:         "/product-carts",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  invalidData3.String(),
			Message:      "you cant donate to yourself. please specify different donor_id and recipient_id"}}

	for _, testcase := range cases {
		req := httptest.NewRequest(testcase.Method, "/", strings.NewReader(testcase.RequestBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("userId", fmt.Sprintf("%v", donorID))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, UpdateProductCartByUserIdController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)

			// data, _ := userResponse.Data.(map[string]interface{})

			// packages, _ := data["recipients"].([]interface{})

			// assert.Equal(t, testcase.Size, len(packages))

		}
	}

}

func Test_DeleteProductCartByUserIdController(t *testing.T) {
	e := InitProductCartTest()

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

	cart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 1, Quantity: 2},
		{RecipientID: uint(children1ID), ProductPackageID: 3, Quantity: 5}, {RecipientID: uint(children2ID), ProductPackageID: 2, Quantity: 1}}

	config.InsertProductCart(cart, donorID)

	validDel := []models.ProductCartDelAPI{{RecipientID: uint(children1ID), ProductPackageID: 3}}

	var validData bytes.Buffer
	json.NewEncoder(&validData).Encode(validDel)

	invalidPackageID := []models.ProductCartDelAPI{{RecipientID: uint(children1ID), ProductPackageID: 19}}

	var invalidData1 bytes.Buffer
	json.NewEncoder(&invalidData1).Encode(invalidPackageID)

	emptyBody := []models.ProductCartDelAPI{}

	var invalidData2 bytes.Buffer
	json.NewEncoder(&invalidData2).Encode(emptyBody)

	cases := []models.UserCaseWithBody{
		{
			Name:         "Delete Item with valid product package id",
			Method:       "DELETE",
			Path:         "/product-carts",
			ExpectedCode: http.StatusOK,
			RequestBody:  validData.String(),
			Message:      "product cart is updated!"},
		{
			Name:         "Delete Item with invalid product package id",
			Method:       "DELETE",
			Path:         "/product-carts",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  invalidData1.String(),
			Message:      "no product_package_id with id: 19 associated with recipient_id 2 found in user's product_carts"},
		{
			Name:         "Delete item without specifying anything",
			Method:       "DELETE",
			Path:         "/product-carts",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  invalidData2.String(),
			Message:      "no item found in delete list. Please specify before deleting"}}

	for _, testcase := range cases {
		req := httptest.NewRequest(testcase.Method, "/", strings.NewReader(testcase.RequestBody))
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("userId", fmt.Sprintf("%v", donorID))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, DeleteProductCartByUserIdController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)

		}
	}

}
