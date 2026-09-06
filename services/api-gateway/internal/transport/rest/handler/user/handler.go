package user_handler

import (
	"api_gateway_microservice/internal/domain"
	user_service "api_gateway_microservice/internal/service/user"
	"api_gateway_microservice/pkg/logger"
	"context"
)

type Handler struct {
	log         *logger.Logger
	userService UserService
}

type UserService interface {
	GetUserById(ctx context.Context, id int64) (domain.User, error)
}

func New(
	log *logger.Logger,
	userService *user_service.Service,
) *Handler {
	return &Handler{
		log:         log,
		userService: userService,
	}
}
