package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func CertificateRoutes() {
	e.GET("/certificates/:completionId", handler.GetCertificateController, middlewares.AuthenticateUser)
}
