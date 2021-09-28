package routes

import (
	handler "berbagi/controllers"
)

func ProficiencyRoute() {
	e.POST("/proficiency", handler.CreateNewProficiencyController)
	e.GET("/proficiency", handler.GetAllProficienciesController)
	e.DELETE("/proficiency/:id", handler.DeleteProficiencyController)
}
