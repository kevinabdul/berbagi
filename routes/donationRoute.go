package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func DonationRoutes() {
	e.GET("/donation", handler.GetDonationsListController, middlewares.AuthenticateUser)
	e.POST("/donation", handler.MakeDonationController, middlewares.AuthenticateUser)
	e.GET("/cart/donation", handler.GetDonationListInCartController, middlewares.AuthenticateUser)
	e.PUT("/cart/donation", handler.UpdateDonationInCartController, middlewares.AuthenticateUser)
	e.DELETE("/cart/donation", handler.DeleteDonationFromCartController, middlewares.AuthenticateUser)
	e.POST("/donation/checkout", handler.CheckoutDonationController, middlewares.AuthenticateUser)
}