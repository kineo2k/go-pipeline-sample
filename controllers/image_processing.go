package controllers

import (
	"encoding/base64"
	"github.com/labstack/echo/v4"
	"go-pipeline-sample/service/pipeline"
	spec2 "go-pipeline-sample/service/spec"
	"io/ioutil"
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

	ticket := pipeline.GetInstance().Enqueue(spec)
	result := <-ticket

	data, err := ioutil.ReadFile(result.OutputPath)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	response := base64.StdEncoding.EncodeToString(data)
	return c.Blob(http.StatusOK, "text/plain", []byte(response))
}
