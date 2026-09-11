package account_grpc_client

import (
	"context"
	"fmt"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
)

func (c *Client) Login(ctx context.Context, email string, password string) (*v1.LoginResponse, error) {
	resp, err := c.api.Login(ctx, &v1.LoginRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, fmt.Errorf("AccountServiceGrpcClient - Login - c.api.Login: %w", err)
	}

	return resp, nil
}
