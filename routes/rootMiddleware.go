package routes

import (
	"github.com/labstack/echo/v4/middleware"
)

func registerRootMiddlewares(){
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	// placeholder untuk cors middleware
}