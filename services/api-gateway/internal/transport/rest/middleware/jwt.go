package http_middleware

import (
	http_lib "api_gateway_microservice/internal/lib/http"
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	"api_gateway_microservice/pkg/logger"
	"fmt"

	"github.com/labstack/echo/v5"
)

type JwtMiddleware struct {
	log     *logger.Logger
	manager *jwt_lib.JWTManager
}

func NewJwtMiddleware(log *logger.Logger, manager *jwt_lib.JWTManager) *JwtMiddleware {
	return &JwtMiddleware{
		log:     log,
		manager: manager,
	}
}

func (m *JwtMiddleware) Process(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		accessTokenCookie, err := c.Cookie("access_token")
		if err != nil {
			return next(c)
		}

		claims, err := m.manager.ParseToken(accessTokenCookie.Value)
		if err != nil {
			return next(c)
		}

		status, err := m.manager.IsBlacklisted(claims.JTI)
		if err != nil {
			m.log.Debug(fmt.Errorf("http_middleware:jwt_middleware:m.manager.IsBlacklisted: %w", err).Error())
			return next(c)
		}

		if !status {
			http_lib.SetClaims(c, claims)
		}

		return next(c)
	}
}
