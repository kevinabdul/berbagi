package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func VolunteerRoutes() {
	e.GET("/volunteers", handler.GetListVolunteersController, middlewares.AuthenticateUser)
	e.GET("/volunteers/profile", handler.GetVolunteerProfileController, middlewares.AuthenticateUser)
}
