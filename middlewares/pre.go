package middlewares

import (
	"github.com/labstack/echo/v4"
)

func ExtractPathAndMethod(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		path := c.Request().URL.Path
		method := c.Request().Method

		c.Request().Header.Set("path", path)
		c.Request().Header.Set("method", method)

		return next(c)
		}
}