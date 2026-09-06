package account_grpc_client

import (
	"context"
	"curret_user_microservice/internal/domain"
	"fmt"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) GetAccountById(ctx context.Context, accountId int64) (*v1.GetAccountByIdResponse, error) {
	resp, err := c.api.GetAccountById(ctx, &v1.GetAccountByIdRequest{
		Id: accountId,
	})
	if err != nil {
		st, ok := status.FromError(err)

		if !ok {
			return nil, fmt.Errorf("AccountServiceGrpcClient - GetAccountById - c.api.GetAccountById: %w", err)
		}

		switch st.Code() {
		case codes.AlreadyExists:
			return nil, domain.ErrAccountNotFound
		}
	}

	return resp, nil
}
