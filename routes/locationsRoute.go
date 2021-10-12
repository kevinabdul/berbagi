package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func LocationsRoute() {
	e.GET("/nearby/:resource", handler.GetAllNearestRecipientsController, middlewares.AuthenticateUser)
}
