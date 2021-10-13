package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func RequestRoute() {
	e.GET("/request", handler.GetAllRequestListController, middlewares.AuthenticateUser)
	e.GET("/request/:field", handler.IdTypeRedirector, middlewares.AuthenticateUser)
	
	e.POST("/request/gift", handler.RequestGift, middlewares.AuthenticateUser)
	e.POST("/request/donation", handler.RequestDonation, middlewares.AuthenticateUser)
	e.POST("/request/service", handler.RequestService, middlewares.AuthenticateUser)
	
	e.DELETE("/request/:request_id", handler.DeleteRequestController, middlewares.AuthenticateUser)
}
