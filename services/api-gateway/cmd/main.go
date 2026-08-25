package main

import (
	"api_gateway_microservice/internal/handler"
	"api_gateway_microservice/pkg/logger"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()

	log := logger.NewLoggerMust()
	defer log.Close()

	e.Logger = log.Logger

	e.Use(middleware.RequestID())

	handler.RegisterHandlers(e)

	if err := e.Start(":8000"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
