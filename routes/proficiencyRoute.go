package routes

import (
	handler "berbagi/controllers"
	"berbagi/middlewares"
)

func ProficiencyRoute() {
	e.POST("/proficiencies", handler.CreateNewProficiencyController, middlewares.AuthenticateUser)
	e.GET("/proficiencies", handler.GetAllProficienciesController, middlewares.AuthenticateUser)
	e.DELETE("/proficiencies/:id", handler.DeleteProficiencyController, middlewares.AuthenticateUser)
	e.PUT("/proficiencies/:id", handler.UpdatedProficiencyController, middlewares.AuthenticateUser)
}
