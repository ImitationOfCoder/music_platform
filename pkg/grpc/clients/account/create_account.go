package account_grpc_client

import (
	"context"
	"fmt"

	errs "github.com/ImitationOfCoder/music_platform/pkg/grpc/errors"
	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) CreateAccount(ctx context.Context, name string, email string, password string) error {
	_, err := c.api.CreateAccount(ctx, &v1.CreateAccountRequest{
		Name:     name,
		Email:    email,
		Password: password,
	})
	if err != nil {
		st, ok := status.FromError(err)

		if !ok {
			return fmt.Errorf("AccountServiceGrpcClient - CreateAccount - c.api.CreateAccount: %w", err)
		}

		switch st.Code() {
		case codes.AlreadyExists:
			return errs.ErrAccountAlreadyExists
		}
	}

	return nil
}
