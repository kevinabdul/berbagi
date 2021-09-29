package routes

import (
	handler "berbagi/controllers"
)

func VolunteerRoutes() {
	e.GET("/volunteers", handler.GetListVolunteers)
}
