package http

import (
	"github.com/labstack/echo/v4"
	"github.com/swaggo/echo-swagger"
)

func NewSwaggerHttpHandler(e *echo.Echo) {
	e.GET("/docs/*", echoSwagger.WrapHandler)
}
