package current_user_handler

import (
	"api_gateway_microservice/internal/domain"
	http_lib "api_gateway_microservice/internal/lib/http"
	"api_gateway_microservice/internal/transport/rest/response"
	"errors"
	"net/http"

	"github.com/labstack/echo/v5"
)

func (h *Handler) GetCurrentUser(c *echo.Context) error {
	ctx := http_lib.BindAccessTokenToContext(c)

	user, err := h.currentUserService.GetCurrentUser(ctx)
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
