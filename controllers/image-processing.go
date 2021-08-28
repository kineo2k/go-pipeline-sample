package controllers

import (
	"github.com/labstack/echo/v4"
	spec2 "go-pipeline-sample/service/spec"
	"log"
	"net/http"
)

type ImageProcessing struct {
	//spec _Spec
}

func NewImageProcessing() *ImageProcessing {
	return &ImageProcessing{}
}

func (*ImageProcessing) Handle(c echo.Context) error {
	spec := spec2.NewSpec()
	if err := c.Bind(spec); err != nil {
		log.Println(err)
		return c.NoContent(http.StatusBadRequest)
	}

	if err := c.Validate(spec); err != nil {
		log.Println(err)
		return c.NoContent(http.StatusBadRequest)
	}

	return c.NoContent(http.StatusOK)
}
