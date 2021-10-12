package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func ConfirmServiceRoutes() {
	e.Renderer = NewRenderer("./*.html", true)
	e.POST("/services/verification", handler.AddConfirmServiceController, middlewares.AuthenticateUser)
	e.GET("/services/verification/:verificationId", handler.GetConfirmServiceController, middlewares.AuthenticateUser)
	e.GET("/services/display/:verificationId", handler.DisplayConfirmServiceController, middlewares.AuthenticateUser)
}
