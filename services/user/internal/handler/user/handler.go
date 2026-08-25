package user_handler

import (
	"user_microservice/internal/domain"
)

type UserHandler struct {
	userService UserService
}

type UserService interface {
	CreateUser(name string, email string, password string) error
	GetUserById(userId int64) (domain.User, error)
}

func NewHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}
