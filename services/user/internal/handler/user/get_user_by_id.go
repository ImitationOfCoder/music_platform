package user_handler

import (
	"errors"
	"net/http"
	"strconv"
	"user_microservice/internal/domain"
	"user_microservice/internal/handler/user/dto"

	"github.com/labstack/echo/v5"
)

func (h *UserHandler) GetUserByID(c *echo.Context) error {
	userIdStr := c.Param("id")
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Invalid user ID",
		})
	}

	user, err := h.userService.GetUserById(userId)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return c.JSON(http.StatusNotFound, map[string]any{
				"success": false,
				"message": "User not found.",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"success": false,
			"message": "Internal server error.",
		})
	}

	return c.JSON(http.StatusOK, map[string]any{
		"success": true,
		"data": map[string]any{
			"user": dto.UserResponse{
				Id:    user.Id,
				Name:  user.Name,
				Email: user.Email,
			},
		},
	})
}
