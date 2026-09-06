package profile_service_grpc_client

import (
	"account_microservice/internal/domain"
	"context"
	"fmt"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) CreateProfile(ctx context.Context, accountId int64, name string) (*v1.CreateProfileResponse, error) {
	resp, err := c.api.CreateProfile(ctx, &v1.CreateProfileRequest{
		AccountId: accountId,
		Name:      name,
	})
	if err != nil {
		if st, ok := status.FromError(err); ok {
			switch st.Code() {
			case codes.AlreadyExists:
				return nil, domain.ErrUserAlreadyExists
			}
		}

		return nil, fmt.Errorf("ProfileServiceGrpcClient - CreateProfile - c.api.CreateProfile: %w", err)
	}

	return resp, nil
}
