package current_user_service

import (
	"api_gateway_microservice/internal/domain"
	"context"
)

func (s *Service) GetCurrentUser(ctx context.Context) (domain.CurrentUser, error) {
	user, err := s.gRPCClients.CurrentUser.GetCurrentUser(ctx)
	if err != nil {
		return domain.CurrentUser{}, domain.ErrUnauthenticated
	}

	return domain.CurrentUser{
		Id:    user.GetId(),
		Name:  user.GetName(),
		Email: user.GetEmail(),
	}, nil
}
