package controllers

import (
	"berbagi/config"
	"berbagi/models"

	"strings"
	"net/url"
	"fmt"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"

	"github.com/stretchr/testify/assert"
	"github.com/labstack/echo/v4"

)

func InitGiftTest() *echo.Echo{
	config.InitDBTest("provinces", "cities","addresses","users", "donors", "childrens", "payment_methods",
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

func Test_GetGiftsController(t *testing.T) {
	e := InitGiftTest()

	donor := models.RegistrationAPI{
		Name: "abdul",
		Email: "abdul@gmail.com",
		Password: "1234",
		NIK: "123450",
		AddressName: "Rumah Abdul",
		ProvinceID: 1,
		CityID: 1,
		Longitude: "123,111",
		Latitude: "111,76",
		RoleID: 2}

	children1 := models.RegistrationAPI{
		Name: "dara",
		Email: "dara@gmail.com",
		Password: "1234",
		NIK: "123451",
		AddressName: "Rumah Dara",
		ProvinceID: 1,
		CityID: 1,
		Longitude: "125,111",
		Latitude: "114,76",
		RoleID: 4}

	children2 := models.RegistrationAPI{
		Name: "ali",
		Email: "ali@gmail.com",
		Password: "1234",
		NIK: "123452",
		AddressName: "Rumah Ali",
		ProvinceID: 1,
		CityID: 1,
		Longitude: "128,111",
		Latitude: "118,76",
		RoleID: 4}
	
	donorID,_ := config.InsertUser(donor)
	children1ID,_ := config.InsertUser(children1)
	children2ID,_ := config.InsertUser(children2)

	cart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 1, Quantity: 2}, 
	{RecipientID: uint(children1ID), ProductPackageID: 3, Quantity: 5}, {RecipientID: uint(children2ID), ProductPackageID: 2, Quantity: 1}}

	config.InsertProductCart(cart, donorID)

	transaction := config.CheckoutProductCart(models.PaymentMethod{ID: uint(1)}, donorID)

	gifts := models.UserCaseWithBody{
		 	Name : "Get all gifts list - exist",
		 	Method: "GET",
			Path : "/gifts",
			ExpectedCode: http.StatusOK,
			RequestBody: "",
			Message: "gifts are retrieved succesfully",
			Size: 1}
		 	

	reqOut := httptest.NewRequest(gifts.Method, "/", nil)
	reqOut.Header.Set("Content-Type", "application/json")
	reqOut.Header.Set("userId", fmt.Sprintf("%v", children1ID))
	recOut := httptest.NewRecorder()
	cOut := e.NewContext(reqOut, recOut)
	
	cOut.SetPath(gifts.Path)

	if assert.NoError(t, GetGiftsController(cOut)) {
		assert.Equal(t, gifts.ExpectedCode, recOut.Code)

		var userResponseOut models.ResponseOK
		
		if err := json.Unmarshal([]byte(recOut.Body.String()), &userResponseOut); err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, gifts.Message, userResponseOut.Message)

		// userData,_ := userResponseOut.Data.(map[string]interface{})
		// details,_ := userData["details"].([]interface{}) 
		// assert.Equal(t, emptyCarts.Size, len(details))
	}

	config.ResolveOnePayment(transaction)

	cases := []models.UserCaseWithBody {
			{Name : "Get pending gifts list - not exist",
		 	Method: "GET",
			Path : "/gifts?status=pending",
			ExpectedCode: http.StatusBadRequest,
			RequestBody: "",
			Message: "no gifts found",
			Size: 0}}


	for _, testcase := range cases {
		var queryValues string
		temp := strings.Split(testcase.Path, "?")
		if len(temp) == 1 {
			queryValues = ""
		} else {
			queryValues = strings.Split(temp[1], "=")[1]
		}

		q := make(url.Values)
		q.Set("status", queryValues)

		req := httptest.NewRequest(testcase.Method, "/?"+q.Encode(), nil)
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("userId", fmt.Sprintf("%v", children1ID))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		
		c.SetPath(testcase.Path)

		if assert.NoError(t, GetGiftsController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK
			
			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}
			
			assert.Equal(t, testcase.Message, userResponse.Message)
			// userData,_ := userResponse.Data.([]interface{})
			// assert.Equal(t, testcase.Size, len(userData))

		}
	}

}

// func Test_AddPendingPaymentController(t *testing.T) {
// 	e := InitPaymentTest()

// 	donor := models.RegistrationAPI{
// 		Name: "abdul",
// 		Email: "abdul@gmail.com",
// 		Password: "1234",
// 		NIK: "123450",
// 		AddressName: "Rumah Abdul",
// 		ProvinceID: 1,
// 		CityID: 1,
// 		Longitude: "123,111",
// 		Latitude: "111,76",
// 		RoleID: 2}

// 	children1 := models.RegistrationAPI{
// 		Name: "dara",
// 		Email: "dara@gmail.com",
// 		Password: "1234",
// 		NIK: "123451",
// 		AddressName: "Rumah Dara",
// 		ProvinceID: 1,
// 		CityID: 1,
// 		Longitude: "125,111",
// 		Latitude: "114,76",
// 		RoleID: 4}

// 	children2 := models.RegistrationAPI{
// 		Name: "ali",
// 		Email: "ali@gmail.com",
// 		Password: "1234",
// 		NIK: "123452",
// 		AddressName: "Rumah Ali",
// 		ProvinceID: 1,
// 		CityID: 1,
// 		Longitude: "128,111",
// 		Latitude: "118,76",
// 		RoleID: 4}
	
// 	donorID,_ := config.InsertUser(donor)
// 	children1ID,_ := config.InsertUser(children1)
// 	children2ID,_ := config.InsertUser(children2)

	
// 	cart := []models.ProductCart{{RecipientID: uint(children1ID), ProductPackageID: 1, Quantity: 2}, 
// 	{RecipientID: uint(children1ID), ProductPackageID: 3, Quantity: 5}, {RecipientID: uint(children2ID), ProductPackageID: 2, Quantity: 1}}

// 	config.InsertProductCart(cart, donorID)

// 	transaction := config.CheckoutProductCart(models.PaymentMethod{ID: uint(1)}, donorID)

// 	invalidInvoice := models.UserPaymentAPI{InvoiceID: "BERBAGI.12312.121.121.222", PaymentMethodID: uint(1), Total: transaction.Total}	
// 	var invalidData1 bytes.Buffer
// 	json.NewEncoder(&invalidData1).Encode(invalidInvoice)

// 	invalidPaymentMethod := models.UserPaymentAPI{InvoiceID: transaction.InvoiceID, PaymentMethodID: uint(9891), Total: transaction.Total}	
// 	var invalidData2 bytes.Buffer
// 	json.NewEncoder(&invalidData2).Encode(invalidPaymentMethod)

// 	invalidAmount := models.UserPaymentAPI{InvoiceID: transaction.InvoiceID, PaymentMethodID: uint(1), Total: 9}	
// 	var invalidData3 bytes.Buffer
// 	json.NewEncoder(&invalidData3).Encode(invalidAmount)

// 	userPaymentValid := models.UserPaymentAPI{InvoiceID: transaction.InvoiceID, PaymentMethodID: uint(1), Total: transaction.Total}
// 	var validData bytes.Buffer
// 	json.NewEncoder(&validData).Encode(userPaymentValid)

// 	cases := []models.UserCaseWithBody {
// 		 {
// 		 	Name : "Resolve pending payments with invalid invoice",
// 		 	Method: "POST",
// 			Path : "/payments",
// 			ExpectedCode: http.StatusBadRequest,
// 			RequestBody: invalidData1.String(),
// 			Message: "No invoice_id found",
// 			Size: 0},
// 		{
// 		 	Name : "Resolve pending payments with incorrect payment method",
// 		 	Method: "POST",
// 			Path : "/payments",
// 			ExpectedCode: http.StatusBadRequest,
// 			RequestBody: invalidData2.String(),
// 			Message:"Specified payment method doesnt match. Please pay using payment_method_id: 1",
// 			Size: 0},
// 		{
// 		 	Name : "Resolve pending payments with incorrect amount",
// 		 	Method: "POST",
// 			Path : "/payments",
// 			ExpectedCode: http.StatusBadRequest,
// 			RequestBody: invalidData3.String(),
// 			Message:fmt.Sprintf("Specified total amount doesnt match. Total mount to be paid: %v", transaction.Total),
// 			Size: 0},
// 		{
// 		 	Name : "Resolve pending payments",
// 		 	Method: "POST",
// 			Path : "/payments",
// 			ExpectedCode: http.StatusOK,
// 			RequestBody: validData.String(),
// 			Message:"payment is succesfull",
// 			Size: 0},
// 		{
// 		 	Name : "Resolve pending payments that already been resolved",
// 		 	Method: "POST",
// 			Path : "/payments",
// 			ExpectedCode: http.StatusBadRequest,
// 			RequestBody: validData.String(),
// 			Message:fmt.Sprintf("Specified invoice id: %v has been paid", transaction.InvoiceID),
// 			Size: 0}}


// 	for _, testcase := range cases {
// 		req := httptest.NewRequest(testcase.Method, "/", strings.NewReader(testcase.RequestBody))
// 		req.Header.Set("Content-Type", "application/json")
// 		req.Header.Set("userId", fmt.Sprintf("%v", donorID))
// 		rec := httptest.NewRecorder()
// 		c := e.NewContext(req, rec)
		
// 		c.SetPath(testcase.Path)

// 		if assert.NoError(t, AddPendingPaymentController(c)) {
// 			assert.Equal(t, testcase.ExpectedCode, rec.Code)

// 			var userResponse models.ResponseOK
			
// 			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
// 				assert.Error(t, err, "error")
// 			}
			
// 			assert.Equal(t, testcase.Message, userResponse.Message)
// 			// userData,_ := userResponse.Data.([]interface{})
// 			// assert.Equal(t, testcase.Size, len(userData))

// 		}
// 	}
// }