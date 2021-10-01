package routes

import (
	"github.com/labstack/echo/v4/middleware"
	"berbagi/middlewares"
)

func rootMiddlewares(){
	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	e.Use(middlewares.ExtractPathAndMethod)

	// placeholder untuk cors middleware
}