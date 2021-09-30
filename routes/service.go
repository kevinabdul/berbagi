package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func ServiceRoute() {
	e.POST("/services", handler.AddServiceToCartController, middlewares.AuthenticateUser)
	e.DELETE("/services", handler.DeleteServiceCartController, middlewares.AuthenticateUser)
	e.PUT("/services", handler.UpdatedServiceOncartController, middlewares.AuthenticateUser)
	e.GET("/services", handler.GetServiceOnCartController, middlewares.AuthenticateUser)
}
