package user_handler

import (
	"net/http"

	"github.com/labstack/echo/v5"
)

func GetUserById(c *echo.Context) error {
	id := c.Param("id")

	return c.JSON(http.StatusOK, id)
}
