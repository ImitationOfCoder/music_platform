package auth_handler

import (
	http_lib "api_gateway_microservice/internal/lib/http"
	"api_gateway_microservice/internal/transport/rest/response"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) Logout(c *echo.Context) error {
	claims := http_lib.GetClaims(c)
	if claims == nil {
		return http_lib.UnauthenticatedResponse(c)
	}

	ctx := c.Request().Context()
	err := h.authService.Logout(ctx, claims)
	if err != nil {
		return http_lib.InternalServerErrorResponse(c)
	}

	http_lib.RemoveSessionCookie(c)

	return c.JSON(http.StatusOK, response.StatusResponse{
		Success: true,
	})
}
