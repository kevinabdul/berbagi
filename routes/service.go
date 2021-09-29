package routes

import (
	handler "berbagi/controllers"
)

func ServiceRoute() {
	e.POST("/service/:id", handler.AddServiceToCartController)
	e.DELETE("/service/:id", handler.DeleteServiceCartController)
}
