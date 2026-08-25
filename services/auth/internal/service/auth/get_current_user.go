package auth_service

import (
	"context"
	"fmt"

	"auth_microservice/internal/domain"
)

func (s *AuthService) GetCurrentUser(ctx context.Context, sessionToken string) (domain.User, error) {
	user, err := s.userRepository.GetUserBySessionToken(ctx, sessionToken)
	if err != nil {
		return domain.User{}, fmt.Errorf("AuthService - GetCurrentUser - s.userRepository.GetUserBySessionToken: %w", err)
	}

	return user, nil
}
