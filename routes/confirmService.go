package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func ConfirmServiceRoutes() {
	e.POST("/services/verification", handler.AddConfirmServiceController, middlewares.AuthenticateUser)
	e.GET("/services/verification", handler.GetConfirmServiceController, middlewares.AuthenticateUser)
}
