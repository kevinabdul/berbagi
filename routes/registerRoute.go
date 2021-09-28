package routes

import(
	handler "berbagi/controllers"
	//"berbagi/middlewares"
	//"fmt"
	//"reflect"
)

func registerRoutes() {
	e.POST("/register/:role", handler.RegisterUserController)

	// e.POST("/register/volunteers", handler.RegisterVolunteerController)

	// e.POST("/register/foundations", handler.RegisterFoundationController)

	// e.POST("/register/childrens", handler.RegisterChildrenController)
}

