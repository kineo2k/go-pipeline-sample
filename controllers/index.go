package controllers

import (
	"github.com/labstack/echo/v4"
	"go-pipeline-sample/tm"
	"net/http"
)

type Index struct {
}

func NewIndex() *Index {
	return &Index{}
}

func (*Index) Handle(c echo.Context) error {
	blob, contentType := tm.GetInstance().GetFile("templates/index.html")
	return c.Blob(http.StatusOK, contentType, blob)
}
