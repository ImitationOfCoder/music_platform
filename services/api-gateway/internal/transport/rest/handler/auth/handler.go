package auth_handler

import (
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	auth_service "api_gateway_microservice/internal/service/auth"
	"api_gateway_microservice/pkg/logger"
	"context"
)

type Handler struct {
	log         *logger.Logger
	authService AuthService
}

type AuthService interface {
	Register(ctx context.Context, name string, email string, password string) error
	Login(ctx context.Context, email string, password string) (string, error)
	Logout(ctx context.Context, claims *jwt_lib.JWTClaims) error
}

func New(
	log *logger.Logger,
	authService *auth_service.Service,
) *Handler {
	return &Handler{
		log:         log,
		authService: authService,
	}
}
