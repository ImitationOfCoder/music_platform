package account_grpc_client

import (
	"fmt"

	"github.com/ImitationOfCoder/music_platform/pkg/logger"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc"
)

type Client struct {
	api v1.AccountServiceClient
	log *logger.Logger
}

func New(log *logger.Logger, addr string, opts []grpc.DialOption) (*Client, error) {
	cc, err := grpc.NewClient(
		addr,
		opts...,
	)
	if err != nil {
		return nil, fmt.Errorf("account_service_grpc_client - New - grpc.NewClient: %w", err)
	}

	return &Client{
		api: v1.NewAccountServiceClient(cc),
		log: log,
	}, nil
}
