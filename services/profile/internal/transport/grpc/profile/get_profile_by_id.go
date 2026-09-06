package profile_grpc

import (
	"context"
	"errors"
	"user_microservice/internal/domain"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) GetProfileById(ctx context.Context, req *v1.GetProfileByIdRequest) (*v1.GetProfileByIdResponse, error) {
	id := req.GetId()

	user, err := s.profileService.GetProfileById(id)
	if err != nil {
		if errors.Is(err, domain.ErrUserNotFound) {
			return nil, status.Error(codes.NotFound, "User not found.")
		}

		return nil, status.Error(codes.Internal, "Internal server error.")
	}

	return &v1.GetProfileByIdResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}
