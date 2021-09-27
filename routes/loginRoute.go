package routes

import(
	handler "berbagi/controllers"
)

func registerLoginRoutes() {
	e.POST("/login", handler.LoginUserController)
}

