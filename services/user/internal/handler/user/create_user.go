package user_handler

import (
	"errors"
	"net/http"
	"user_microservice/internal/domain"
	"user_microservice/internal/handler/user/dto"

	"github.com/labstack/echo/v5"
)

func (h *UserHandler) CreateUser(c *echo.Context) error {
	var body dto.CreateUserRequest
	if errs := c.Bind(&body); errs != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"success": false,
			"message": "Invalid request data.",
		})
	}

	err := h.userService.CreateUser(
		body.Name,
		body.Email,
		body.Password,
	)
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return c.JSON(http.StatusConflict, map[string]any{
				"success": false,
				"message": "User already exists.",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]any{
			"success": false,
			"message": "Internal server error.",
		})
	}

	return c.JSON(http.StatusCreated, map[string]any{
		"success": true,
	})
}
