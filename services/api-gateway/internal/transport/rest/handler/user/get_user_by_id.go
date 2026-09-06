package user_handler

import (
	"api_gateway_microservice/internal/domain"
	http_lib "api_gateway_microservice/internal/lib/http"
	"api_gateway_microservice/internal/transport/rest/response"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) GetUserById(c *echo.Context) error {
	id, ok := http_lib.GetInt64Param(c, "id")
	if !ok {
		return http_lib.UserNotFoundResponse(c)
	}

	ctx := c.Request().Context()

	user, err := h.userService.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return http_lib.UserNotFoundResponse(c)
		}

		return http_lib.InternalServerErrorResponse(c)
	}

	return c.JSON(http.StatusOK, response.BaseResponse{
		Success: true,
		Data: response.UserResponse{
			Id:   user.Id,
			Name: user.Name,
		},
	})
}
