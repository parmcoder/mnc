package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/parmcoder/mnc"
)

type Base interface {
	Create(c echo.Context) error
	HealthCheck(c echo.Context) error
	Monitor(ctx echo.Context) error
}

type BaseImpl struct {
	apiConnector *mnc.Base
}

func CreateController(ac *mnc.Base) Base {
	return BaseImpl{
		apiConnector: ac,
	}
}

func (c BaseImpl) Create(ctx echo.Context) error {
	var message mnc.CreateBroadcast

	err := ctx.Bind(&message)
	if err != nil {
		return ctx.String(http.StatusBadRequest, err.Error())
	}

	response, err := (*c.apiConnector).Broadcast(&message)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}
	return ctx.JSON(http.StatusOK, response)
}

func (c BaseImpl) Monitor(ctx echo.Context) error {
	var message mnc.GetTransaction

	hashMessage := ctx.Param("hash")
	message.TxHash = hashMessage

	response, err := (*c.apiConnector).Monitor(&message)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, response)
}

func (c BaseImpl) HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "Server started!")
}
