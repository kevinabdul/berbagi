package routes

import(
	handler "berbagi/controllers"
)

func registerLoginRoutes() {
	e.POST("/register/personal", handler.RegisterPersonalRecipientController)
	e.POST("/register/agency", handler.RegisterAgencyRecipientController)
	e.POST("/register/donator", handler.RegisterDonatorController)
	e.POST("/register/volunteer", handler.RegisterVolunteerController)
	
	e.POST("/login", handler.LoginUserController)
}

