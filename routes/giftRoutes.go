package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func giftRoutes() {
	e.GET("/gifts", handler.GetGiftsController, middlewares.AuthenticateUser, middlewares.AuthorizeUser)
}