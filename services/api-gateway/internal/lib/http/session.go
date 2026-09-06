package http_lib

import (
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v5"
	"google.golang.org/grpc/metadata"
)

func SetClaims(c *echo.Context, claims *jwt_lib.JWTClaims) {
	c.Set("user", claims)
}

func GetClaims(c *echo.Context) *jwt_lib.JWTClaims {
	claims, err := echo.ContextGet[*jwt_lib.JWTClaims](c, "user")
	if err != nil {
		return nil
	}

	return claims
}

func BindAccessTokenToContext(c *echo.Context) context.Context {
	cookie, err := c.Cookie("access_token")
	if err != nil {
		return nil
	}

	ctx := c.Request().Context()
	mt := metadata.New(map[string]string{
		"authorization": fmt.Sprintf("Bearer %s", cookie.Value),
	})

	return metadata.NewOutgoingContext(ctx, mt)
}

func SetSessionCookie(c *echo.Context, token string) {
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   int(24 * time.Hour * 30),
	})
}

func RemoveSessionCookie(c *echo.Context) {
	c.SetCookie(&http.Cookie{
		Name:     "access_token",
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		MaxAge:   -1,
	})
}
