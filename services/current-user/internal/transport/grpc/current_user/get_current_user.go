package current_user_grpc

import (
	"context"
	"strings"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) GetCurrentUserProfile(ctx context.Context, req *v1.GetCurrentUserProfileRequest) (*v1.GetCurrentUserProfileResponse, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.InvalidArgument, "missing metadata")
	}

	authHeader := md["authorization"]
	if len(authHeader) != 1 {
		return nil, status.Error(codes.InvalidArgument, "missing authorization header")
	}

	token := authHeader[0]
	token = strings.TrimPrefix(token, "Bearer ")
	token = strings.TrimSpace(token)

	return &v1.GetCurrentUserProfileResponse{
		Id:    1,
		Name:  token,
		Email: "test",
	}, nil
}
