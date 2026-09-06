package auth_handler

import (
	"api_gateway_microservice/internal/domain"
	http_lib "api_gateway_microservice/internal/lib/http"
	"api_gateway_microservice/internal/transport/rest/request"
	"api_gateway_microservice/internal/transport/rest/response"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) Login(c *echo.Context) error {
	claims := http_lib.GetClaims(c)
	if claims != nil {
		return http_lib.AlreadyAuthenticatedResponse(c)
	}

	var body request.LoginRequest
	if err := c.Bind(&body); err != nil {
		return http_lib.BadRequestResponse(c)
	}

	ctx := c.Request().Context()
	token, err := h.authService.Login(ctx, body.Email, body.Password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return http_lib.InvalidCredentialsResponse(c)
		}

		return http_lib.InternalServerErrorResponse(c)
	}

	http_lib.SetSessionCookie(c, token)

	return c.JSON(http.StatusOK, response.StatusResponse{
		Success: true,
	})
}
