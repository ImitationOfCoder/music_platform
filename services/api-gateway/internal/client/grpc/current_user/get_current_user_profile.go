package current_user_grpc_client

import (
	"api_gateway_microservice/internal/domain"
	"context"
	"fmt"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) GetCurrentUserProfile(ctx context.Context) (domain.CurrentUser, error) {
	resp, err := c.api.GetCurrentUserProfile(ctx, &v1.GetCurrentUserProfileRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return domain.CurrentUser{}, fmt.Errorf("CurrentUserServiceGrpcClient - GetCurrentUserProfile - c.api.GetCurrentUserProfile: %w", err)
		}

		switch st.Code() {
		case codes.Unauthenticated:
			return domain.CurrentUser{}, domain.ErrUnauthenticated
		}
	}

	return domain.CurrentUser{
		Id:    resp.GetId(),
		Name:  resp.GetName(),
		Email: resp.GetEmail(),
	}, nil
}
