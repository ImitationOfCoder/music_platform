package current_user_handler

import (
	"api_gateway_microservice/internal/domain"
	http_lib "api_gateway_microservice/internal/lib/http"
	"api_gateway_microservice/internal/transport/rest/response"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) UpdateCurrentUser(c *echo.Context) error {
	claims := http_lib.GetClaims(c)
	if claims == nil {
		return http_lib.UnauthenticatedResponse(c)
	}

	name := c.FormValue("name")

	ctx := c.Request().Context()

	user, err := h.currentUserService.UpdateCurrentUser(ctx, claims, name)
	if err != nil {
		if errors.Is(err, domain.ErrUnauthenticated) {
			return http_lib.UnauthenticatedResponse(c)
		}

		return http_lib.InternalServerErrorResponse(c)
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Success: true,
		Data: response.CurrentUserResponse{
			Id:    user.Id,
			Name:  user.Name,
			Email: user.Email,
		},
	})
}
