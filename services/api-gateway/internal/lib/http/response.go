package http_lib

import (
	"api_gateway_microservice/internal/transport/rest/response"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v5"
)

func InternalServerErrorResponse(c *echo.Context) error {
	return c.JSON(http.StatusInternalServerError, response.BaseResponse{
		Success: false,
		Data: map[string]string{
			"message": "Internal Server Error.",
		},
	})
}

func UnauthenticatedResponse(c *echo.Context) error {
	return c.JSON(http.StatusUnauthorized, response.BaseResponse{
		Success: false,
		Data: map[string]string{
			"message": "Unauthenticated.",
		},
	})
}

func InvalidCredentialsResponse(c *echo.Context) error {
	return c.JSON(http.StatusUnauthorized, response.BaseResponse{
		Success: false,
		Data: map[string]string{
			"message": "Invalid credentials.",
		},
	})
}

func AlreadyAuthenticatedResponse(c *echo.Context) error {
	return c.JSON(http.StatusConflict, response.BaseResponse{
		Success: false,
		Data: map[string]string{
			"message": "You are already logged in.",
		},
	})
}

func UserNotFoundResponse(c *echo.Context) error {
	return c.JSON(http.StatusNotFound, response.BaseResponse{
		Success: false,
		Data: map[string]string{
			"message": "User not found.",
		},
	})
}

func EntityAlreadyExistsResponse(c *echo.Context, entity string) error {
	return c.JSON(http.StatusConflict, response.BaseResponse{
		Success: false,
		Data: map[string]string{
			"message": fmt.Sprintf("%s already exists.", entity),
		},
	})
}

func BadRequestResponse(c *echo.Context) error {
	return c.JSON(http.StatusBadRequest, map[string]any{
		"success": false,
		"message": "Invalid request body.",
	})
}
