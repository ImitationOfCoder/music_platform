package auth_service

import (
	"api_gateway_microservice/internal/domain"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) Login(ctx context.Context, email string, password string) (string, error) {
	account, err := s.gRPCClients.Account.Login(
		ctx,
		email,
		password,
	)
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return "", fmt.Errorf("AuthService - Login - s.gRPCClients.Account.Login: %w", err)
		}

		switch st.Code() {
		case codes.Unauthenticated:
			return "", domain.ErrInvalidCredentials
		default:
			return "", fmt.Errorf("AuthService - Login - s.gRPCClients.Account.Login: %w", st.Err())
		}
	}

	token, err := s.jwtManager.NewToken(
		account.GetId(),
		account.GetEmail(),
	)
	if err != nil {
		return "", fmt.Errorf("AuthService - Login - s.jwtManager.NewToken: %w", err)
	}

	return token, nil
}
