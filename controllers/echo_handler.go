package controllers

import "github.com/labstack/echo/v4"

type EchoHandler interface {
	Handle(c echo.Context) error
}
