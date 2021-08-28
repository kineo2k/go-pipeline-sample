package controllers

import (
	"github.com/labstack/echo/v4"
	"go-pipeline-sample/tm"
	"net/http"
)

type Statics struct {
}

func NewStatics() *Statics {
	return &Statics{}
}

func (*Statics) Handle(c echo.Context) error {
	path := c.Request().RequestURI

	if tm.GetInstance().Exists(path) {
		blob, contentType := tm.GetInstance().GetFile(path)
		return c.Blob(http.StatusOK, contentType, blob)
	}

	return c.NoContent(http.StatusNotFound)
}
