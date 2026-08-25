package auth_service

import (
	"context"

	"auth_microservice/internal/domain"
)

type AuthService struct {
	userRepository    UserRepository
	sessionRepository SessionRepository
}

type UserRepository interface {
	Login(
		ctx context.Context,
		email string,
		password string,
	) (int64, error)
	GetUserBySessionToken(
		ctx context.Context,
		sessionToken string,
	) (domain.User, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, userId int64) (string, error)
	DeleteSession(ctx context.Context, token string) error
}

func NewService(
	userRepository UserRepository,
	sessionRepository SessionRepository,
) *AuthService {
	return &AuthService{
		userRepository:    userRepository,
		sessionRepository: sessionRepository,
	}
}
