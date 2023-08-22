package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/parmcoder/transaction-tracker/controllers"
)

func main() {
	fmt.Println("hello world!")

	e := echo.New()

	e.GET("/", controllers.HealthCheck)

	e.Logger.Fatal(e.Start(":1323"))
}
