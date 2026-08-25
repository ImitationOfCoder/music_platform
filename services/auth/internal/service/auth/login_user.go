package auth_service

import (
	"context"
	"errors"
	"fmt"

	"auth_microservice/internal/domain"
)

func (s *AuthService) Login(ctx context.Context, email, password string) (string, error) {
	id, err := s.userRepository.Login(ctx, email, password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return "", domain.ErrInvalidCredentials
		}

		return "", fmt.Errorf("AuthService - Login - s.userRepository.Login: %w", err)
	}

	token, err := s.sessionRepository.CreateSession(ctx, id)
	if err != nil {
		return "", fmt.Errorf("AuthService - Login - s.sessionRepository.CreateSession: %w", err)
	}

	return token, nil
}
