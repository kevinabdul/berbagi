package routes

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func New() *echo.Echo {
	e = echo.New()

	rootMiddlewares()

	registerRoutes()

	loginRoutes()

	ProficiencyRoute()

	productCartRoutes()
	
	return e
}