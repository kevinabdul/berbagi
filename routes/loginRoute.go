package routes

import(
	handler "berbagi/controllers"
)

func loginRoutes() {
	e.POST("/login", handler.LoginUserController)
}

