package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func LocationsRoute() {
	e.GET("/nearbyrecipient", handler.GetAllNearestRecipientsController, middlewares.AuthenticateUser)
}
