package http_lib

import (
	"strconv"

	"github.com/labstack/echo/v5"
)

func GetInt64Param(c *echo.Context, value string) (int64, bool) {
	val, err := strconv.ParseInt(c.Param(value), 10, 64)
	if err != nil {
		return 0, false
	}

	return val, true
}
