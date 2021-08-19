package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func EchoStart(port int32) {
	e := echo.New()

	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "<h1>Go Pipeline Pattern Sample</h1>")
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
