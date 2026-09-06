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

func (s *ServerApi) CreateAccount(ctx context.Context, req *v1.CreateAccountRequest) (*v1.CreateAccountResponse, error) {
	err := validateCreateAccount(req)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "validation failed: %s", err.Error())
	}

	name := req.GetName()
	email := req.GetEmail()
	password := req.GetPassword()

	err = s.accountService.CreateAccount(ctx, name, email, password)
	if err != nil {
		if errors.Is(err, domain.ErrAccountAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "User already exists.")
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Server Error: %s", err.Error()))
	}

	return &v1.CreateAccountResponse{
		Success: true,
	}, nil
}

func validateCreateAccount(req *v1.CreateAccountRequest) error {
	name := req.GetName()

	if len([]rune(name)) < 1 {
		return fmt.Errorf("user name cannot be empty")
	}

	if len([]rune(name)) > 32 {
		return fmt.Errorf("user name length cant be greater than 32 characters")
	}

	email := req.GetEmail()

	if len([]rune(email)) < 1 {
		return fmt.Errorf("email cannot be empty")
	}

	password := req.GetPassword()

	if len([]rune(password)) < 1 {
		return fmt.Errorf("password is required")
	}

	return nil
}
