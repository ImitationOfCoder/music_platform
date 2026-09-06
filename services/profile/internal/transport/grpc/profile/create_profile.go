package profile_grpc

import (
	"context"
	"errors"
	"fmt"
	"user_microservice/internal/domain"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) CreateProfile(ctx context.Context, req *v1.CreateProfileRequest) (*v1.CreateProfileResponse, error) {
	err := validateCreateUser(req)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	err = s.profileService.CreateProfile(
		req.GetAccountId(),
		req.GetName(),
	)
	if err != nil {
		if errors.Is(err, domain.ErrUserAlreadyExists) {
			return nil, status.Error(codes.AlreadyExists, "User already exists.")
		}

		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &v1.CreateProfileResponse{
		Success: true,
	}, nil
}

func validateCreateUser(req *v1.CreateProfileRequest) error {
	name := req.GetName()

	if len([]rune(name)) < 1 {
		return fmt.Errorf("user name cannot be empty")
	}

	if len([]rune(name)) > 32 {
		return fmt.Errorf("user name length cant be greater than 32 characters")
	}

	return nil
}
