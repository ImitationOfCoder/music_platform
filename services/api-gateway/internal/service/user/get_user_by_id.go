package user_service

import (
	"api_gateway_microservice/internal/domain"
	"context"
	"errors"
	"fmt"
)

func (s *Service) GetUserById(ctx context.Context, id int64) (domain.User, error) {
	profile, err := s.gRPCClients.Profile.GetProfileById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrProfileNotFound) {
			return domain.User{}, domain.ErrUserNotFound
		}

		return domain.User{}, fmt.Errorf("UserService - GetUserById - s.gRPCClients.Profile.GetProfileById: %w", err)
	}

	return domain.User{
		Id:   profile.Id,
		Name: profile.Name,
	}, nil
}
