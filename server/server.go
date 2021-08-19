package server

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-pipeline-sample/tm"
	"net/http"
)

func EchoStart(port int32) {
	e := echo.New()

	e.Use(middleware.Recover())

	// Statics
	e.GET("/*", func(c echo.Context) error {
		path := c.Request().RequestURI

		if tm.GetInstance().Exists(path) {
			blob, contentType := tm.GetInstance().GetFile(path)
			return c.Blob(http.StatusOK, contentType, blob)
		}

		return c.NoContent(http.StatusNotFound)
	})

	// Index
	e.GET("/", func(c echo.Context) error {
		blob, contentType := tm.GetInstance().GetFile("templates/index.html")
		return c.Blob(http.StatusOK, contentType, blob)
	})

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
