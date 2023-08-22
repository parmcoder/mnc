package core

import (
	"github.com/labstack/echo/v4"
)

func Setup() *echo.Echo {
	e := echo.New()

	e.GET("/", c.HealthCheck)
	e.GET("monitor/:hash", c.Monitor)
	e.POST("/", c.Create)

	return e
}
