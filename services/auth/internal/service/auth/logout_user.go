package auth_service

import (
	"context"
)

func (s *AuthService) LogOutUser(ctx context.Context, token string) error {
	return s.sessionRepository.DeleteSession(ctx, token)
}
