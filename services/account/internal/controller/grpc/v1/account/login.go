package account_grpc

import (
	"account_microservice/internal/domain"
	"context"
	"errors"
	"fmt"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginResponse, error) {
	email := req.GetEmail()
	password := req.GetPassword()

	account, err := s.accountService.Login(ctx, email, password)
	if err != nil {
		if errors.Is(err, domain.ErrInvalidCredentials) {
			return nil, status.Error(codes.Unauthenticated, fmt.Sprintf("Invalid credentials."))
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Server Error: %s", err.Error()))
	}

	return &v1.LoginResponse{
		Id:    account.Id,
		Email: account.Email,
	}, nil
}
