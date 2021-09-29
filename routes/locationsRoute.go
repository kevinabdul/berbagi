package routes

import (
	handler "berbagi/controllers"
)

func LocationsRoute() {
	e.GET("/nearbyrecipient", handler.GetAllNearestRecipientsController)
}
