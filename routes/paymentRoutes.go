package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func paymentRoutes() {
	e.GET("/payments", handler.GetPendingPaymentsController, middlewares.AuthenticateUser)

	e.POST("/payments", handler.AddPendingPaymentController, middlewares.AuthenticateUser)
}