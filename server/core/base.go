package core

import (
	"github.com/labstack/echo/v4"
	"github.com/parmcoder/mnc"
	"github.com/parmcoder/transaction-tracker/controllers"
)

func Setup() *echo.Echo {
	e := echo.New()

	// This is how you can use the service to inject the dependency
	ac := mnc.CreateAPIConnector()
	c := controllers.CreateController(&ac)

	e.GET("health", c.HealthCheck)
	e.GET("mnc/:hash", c.Monitor)
	e.POST("mnc", c.Create)

	return e
}
