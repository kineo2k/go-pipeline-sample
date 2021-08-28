package server

import (
	"fmt"
	validator2 "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-pipeline-sample/controllers"
)

func EchoStart(port int32) {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Validator = &RestApiValidator{validator: validator2.New()}

	// Routing Rules
	e.GET("/", controllers.NewIndex().Handle)
	e.GET("/*", controllers.NewStatics().Handle)
	e.POST("/image-processing", controllers.NewImageProcessing().Handle)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))
}
