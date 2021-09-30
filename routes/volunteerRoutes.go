package routes

import (
	handler "berbagi/controllers"
)

func VolunteerRoutes() {
	e.GET("/volunteers", handler.GetListVolunteers)
	e.GET("/volunteers/:id", handler.GetVolunteerProfileController)
}
