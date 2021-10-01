package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func CompletionRoutes() {
	e.GET("/completion/:verificationId", handler.GetCompletionDetailController, middlewares.AuthenticateUser)
	e.PUT("/completion/:verificationId", handler.UpdateStatusCompletionController, middlewares.AuthenticateUser)
}
