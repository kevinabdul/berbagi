package controllers

import (
	libdb "berbagi/lib/database"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCertificateController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	completionId, errorId := strconv.Atoi(c.Param("completionId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid completion id"})
	}

	certificate, rowAffected, err := libdb.GetCertificateService(volunteerId, completionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get certificate"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "completion id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorize access"})
	}

	return c.JSON(http.StatusOK, struct {
		Status  string
		Message string
		Data    interface{}
	}{Status: "Success", Message: "success to get certificate", Data: certificate})

}

type Renderer struct {
	template *template.Template
	debug    bool
	location string
}

func NewRenderer(location string, debug bool) *Renderer {
	tpl := new(Renderer)
	tpl.location = location
	tpl.debug = debug

	tpl.ReloadTemplates()

	return tpl
}

func (t *Renderer) ReloadTemplates() {
	t.template = template.Must(template.ParseGlob(t.location))
}

func (t *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if t.debug {
		t.ReloadTemplates()
	}

	return t.template.ExecuteTemplate(w, name, data)
}

func CertificateDisplayController(c echo.Context) error {
	volunteerId, _ := strconv.Atoi(c.Request().Header.Get("userId"))

	completionId, errorId := strconv.Atoi(c.Param("completionId"))
	if errorId != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Invalid completion id"})
	}

	certificate, rowAffected, err := libdb.GetCertificateService(volunteerId, completionId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Failed to get certificate"})
	}

	if rowAffected == 0 {
		return c.JSON(http.StatusOK, struct {
			Status  string
			Message string
		}{Status: "Success", Message: "completion id not found"})
	}

	if rowAffected == -1 {
		return c.JSON(http.StatusBadRequest, struct {
			Status  string
			Message string
		}{Status: "Failed", Message: "Unauthorize access"})
	}

	return c.Render(http.StatusOK, "index.html", certificate)
}
