package profile_grpc_client

import (
	"context"
	"fmt"

	errs "github.com/ImitationOfCoder/music_platform/pkg/grpc/errors"
	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (c *Client) GetProfileById(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error) {
	resp, err := c.api.GetProfileById(ctx, &v1.GetProfileByIdRequest{
		Id: id,
	})
	if err != nil {
		st, ok := status.FromError(err)
		if !ok {
			return nil, fmt.Errorf("ProfileServiceGrpcClient - GetProfileById - c.api.GetProfileById: %w", err)
		}

		switch st.Code() {
		case codes.NotFound:
			return nil, errs.ErrProfileNotFound
		}
	}

	return resp, nil
}
