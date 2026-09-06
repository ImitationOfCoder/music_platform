package current_user_grpc_client

import (
	"api_gateway_microservice/pkg/logger"
	"fmt"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc"
)

type Client struct {
	api v1.CurrentUserServiceClient
	log *logger.Logger
}

func New(log *logger.Logger, addr string, opts []grpc.DialOption) (*Client, error) {
	cc, err := grpc.NewClient(
		addr,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("profile_service_grpc_client - New - grpc.NewClient: %w", err)
	}

	return &Client{
		api: v1.NewCurrentUserServiceClient(cc),
		log: log,
	}, nil
}
