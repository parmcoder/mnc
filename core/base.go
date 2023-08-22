package core

import (
	"github.com/labstack/echo/v4"
	"github.com/parmcoder/transaction-tracker/controllers"
	"github.com/parmcoder/transaction-tracker/services"
)

func Setup() *echo.Echo {
	e := echo.New()
	ac := services.CreateAPIConnector()
	c := controllers.CreateController(&ac)

	e.GET("/", c.HealthCheck)
	e.GET("monitor/:hash", c.Monitor)
	e.POST("/", c.Create)

	return e
}
