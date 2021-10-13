package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
)

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

func CertificateRoutes() {
	e.Renderer = NewRenderer("./*.html", true)
	e.GET("/certificates/:completionId", handler.GetCertificateController, middlewares.AuthenticateUser)
	e.GET("/certificates/display/:completionId", handler.CertificateDisplayController, middlewares.AuthenticateUser)
}
