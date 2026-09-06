package auth_service

import (
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	"context"
	"fmt"
)

func (s *Service) Logout(ctx context.Context, claims *jwt_lib.JWTClaims) error {
	err := s.jwtManager.AddToBlacklist(claims)
	if err != nil {
		return fmt.Errorf("AuthService - Logout - s.jwtManager.AddToBlacklist: %w", err)
	}

	return nil
}
