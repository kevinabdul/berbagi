package routes

import (
	handler "berbagi/controllers"
)

func ProficiencyRoute() {
	e.POST("/proficiency", handler.CreateNewProficiencyController)

}
