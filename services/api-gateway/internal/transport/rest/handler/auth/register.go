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

func (h *Handler) Register(c *echo.Context) error {
	claims := http_lib.GetClaims(c)
	if claims != nil {
		return http_lib.AlreadyAuthenticatedResponse(c)
	}

	var body request.RegisterRequest
	if err := c.Bind(&body); err != nil {
		return http_lib.BadRequestResponse(c)
	}

	ctx := c.Request().Context()
	err := h.authService.Register(ctx, body.Name, body.Email, body.Password)
	if err != nil {
		if errors.Is(err, domain.ErrAccountAlreadyExists) {
			return http_lib.EntityAlreadyExistsResponse(c, "Account")
		}

		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": true,
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, response.StatusResponse{
		Success: true,
	})
}
