package current_user_handler

import (
	"api_gateway_microservice/internal/domain"
	current_user_service "api_gateway_microservice/internal/service/current_user"
	"api_gateway_microservice/pkg/logger"
	"context"
)

type Handler struct {
	log                *logger.Logger
	currentUserService CurrentUserService
}

type CurrentUserService interface {
	GetCurrentUser(ctx context.Context) (domain.CurrentUser, error)
}

func New(log *logger.Logger, currentUserService *current_user_service.Service) *Handler {
	return &Handler{
		log:                log,
		currentUserService: currentUserService,
	}
}
