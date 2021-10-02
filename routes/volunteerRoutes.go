package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func VolunteerRoutes() {
	e.GET("/volunteers", handler.GetListVolunteersController, middlewares.AuthenticateUser)
	e.GET("/volunteer", handler.GetVolunteerProfileController, middlewares.AuthenticateUser)
}
