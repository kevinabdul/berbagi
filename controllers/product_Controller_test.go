package controllers

import (
	"berbagi/config"
	"berbagi/models"

	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InitProductTest() *echo.Echo {
	config.InitDBTest("categories", "products")

	e := echo.New()
	return e
}

func Test_GetProductsController(t *testing.T) {
	e := InitProductTest()

	emptyCase := models.UserCaseWithBody{
		Name:         "Get Product from empty table",
		Method:       "GET",
		Path:         "/products",
		ExpectedCode: http.StatusBadRequest,
		RequestBody:  "",
		Message:      "No product found in the product table"}

	var queryValues string
	temp := strings.Split(emptyCase.Path, "?")
	if len(temp) == 1 {
		queryValues = ""
	} else {
		queryValues = strings.Split(temp[1], "=")[1]
	}

	q := make(url.Values)
	q.Set("categoryId", queryValues)

	req := httptest.NewRequest("GET", "/?"+q.Encode(), nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(emptyCase.Path)

	if assert.NoError(t, GetProductsController(c)) {
		assert.Equal(t, emptyCase.ExpectedCode, rec.Code)

		var userResponse models.ResponseOK

		if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
			assert.Error(t, err, "error")
		}

		assert.Equal(t, emptyCase.Message, userResponse.Message)
	}

	config.InsertCategory()
	config.InsertProduct()

	cases := []models.UserCaseWithBody{
		{
			Name:         "Get product packages",
			Method:       "GET",
			Path:         "/products",
			ExpectedCode: http.StatusOK,
			RequestBody:  "",
			Message:      "Products retrieval are succesfull"},
		{
			Name:         "Get Product with invalid category id",
			Method:       "GET",
			Path:         "/products?category=45",
			ExpectedCode: http.StatusBadRequest,
			RequestBody:  "",
			Message:      "No product found for the given category"}}

	for _, testcase := range cases {
		var queryValues string
		temp := strings.Split(testcase.Path, "?")
		if len(temp) == 1 {
			queryValues = ""
		} else {
			queryValues = strings.Split(temp[1], "=")[1]
		}

		q := make(url.Values)
		q.Set("categoryId", queryValues)

		req := httptest.NewRequest("GET", "/?"+q.Encode(), nil)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(testcase.Path)

		if assert.NoError(t, GetProductsController(c)) {
			assert.Equal(t, testcase.ExpectedCode, rec.Code)

			var userResponse models.ResponseOK

			if err := json.Unmarshal([]byte(rec.Body.String()), &userResponse); err != nil {
				assert.Error(t, err, "error")
			}

			assert.Equal(t, testcase.Message, userResponse.Message)
		}
	}
}
