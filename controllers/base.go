package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/parmcoder/transaction-tracker/models"
	"github.com/parmcoder/transaction-tracker/services"
)

type Base interface {
	Create(c echo.Context) error
	HealthCheck(c echo.Context) error
	Monitor(ctx echo.Context) error
}

type BaseImpl struct {
	apiConnector *services.Base
}

func CreateController(ac *services.Base) Base {
	return BaseImpl{
		apiConnector: ac,
	}
}

func (c BaseImpl) Create(ctx echo.Context) error {
	message := models.CreateBroadcast{}
	response, err := (*c.apiConnector).Broadcast(&message)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c BaseImpl) Monitor(ctx echo.Context) error {
	message := models.GetTransaction{}
	response, err := (*c.apiConnector).Monitor(&message)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c BaseImpl) HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Server started!")
}
