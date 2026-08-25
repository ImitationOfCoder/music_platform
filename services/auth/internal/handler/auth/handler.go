package auth_handler

import (
	"context"
	"net/http"

	"auth_microservice/internal/domain"

	"github.com/ImitationOfCoder/music_platform/pkg/http_server"
)

type AuthHandler struct {
	authService AuthService
}

type AuthService interface {
	GetCurrentUser(ctx context.Context, token string) (domain.User, error)
	Login(ctx context.Context, email string, password string) (string, error)
	LogOutUser(ctx context.Context, token string) error
}

func NewHandler(authService AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Routes() []http_server.Route {
	return []http_server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/login",
			Handler: h.Login,
		},
		{
			Method:  http.MethodGet,
			Path:    "/auth/me",
			Handler: h.GetCurrentUser,
		},
		{
			Method:  http.MethodDelete,
			Path:    "/auth/logout",
			Handler: h.LogOutUser,
		},
	}
}
