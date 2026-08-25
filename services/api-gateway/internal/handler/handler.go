package handler

import (
	user_handler "api_gateway_microservice/internal/handler/user"

	"github.com/labstack/echo/v5"
)

func RegisterHandlers(e *echo.Echo) {
	e.GET("/api/users/:id", user_handler.GetUserById)
}
