package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func paymentRoutes() {
	e.GET("/payments", handler.GetPendingPaymentsController, middlewares.AuthenticateUser, middlewares.AuthorizeUser)

	e.POST("/payments", handler.AddPendingPaymentController, middlewares.AuthenticateUser, middlewares.AuthorizeUser)
	e.POST("/payments/donation", handler.AddPendingDonationPaymentController, middlewares.AuthenticateUser, middlewares.AuthorizeUser)
}