package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func CertificateRoutes() {
	e.Renderer = handler.NewRenderer("./*.html", true)
	e.GET("/certificates/:completionId", handler.GetCertificateController, middlewares.AuthenticateUser)
	e.GET("/certificates/display/:completionId", handler.CertificateDisplayController, middlewares.AuthenticateUser)
}
