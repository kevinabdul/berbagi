package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func checkoutRoutes() {
	e.GET("/checkout", handler.GetCheckoutByUserIdController, middlewares.AuthenticateUser, middlewares.AuthorizeUser)

	e.POST("/checkout", handler.AddCheckoutByUserIdController, middlewares.AuthenticateUser, middlewares.AuthorizeUser)
}