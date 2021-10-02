package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func RequestRoute() {
	e.POST("/request/gift", handler.RequestGift, middlewares.AuthenticateUser)
	e.POST("/request/donation", handler.RequestDonation, middlewares.AuthenticateUser)
	e.POST("/request/service", handler.RequestService, middlewares.AuthenticateUser)
}
