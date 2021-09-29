package routes

import (
	handler "berbagi/controllers"
	middleware "berbagi/middlewares"
)

func ProficiencyRoute() {
	e.POST("/proficiency", handler.CreateNewProficiencyController)
	e.GET("/proficiency", handler.GetAllProficienciesController, middleware.AuthenticateUser)
	e.DELETE("/proficiency/:id", handler.DeleteProficiencyController)
	e.PUT("/proficiency/:id", handler.UpdatedProficiencyController)
}
