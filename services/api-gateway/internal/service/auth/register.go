package auth_service

import (
	"api_gateway_microservice/internal/domain"
	"context"
	"errors"
	"fmt"
)

func (s *Service) Register(
	ctx context.Context,
	name string,
	email string,
	password string,
) error {
	err := s.gRPCClients.Account.CreateAccount(
		ctx,
		name,
		email,
		password,
	)
	if err != nil {
		if errors.Is(err, domain.ErrAccountAlreadyExists) {
			return domain.ErrAccountAlreadyExists
		}

		return fmt.Errorf("AuthService - Register - s.gRPCClients.Account.CreateAccount: %w", err)
	}

	return nil
}
