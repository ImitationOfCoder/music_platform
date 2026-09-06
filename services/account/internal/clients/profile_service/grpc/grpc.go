package profile_service_grpc_client

import (
	"account_microservice/pkg/logger"
	"context"
	"fmt"
	"time"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Client struct {
	api v1.ProfileServiceClient
	log *logger.Logger
}

func New(ctx context.Context, log *logger.Logger, addr string, timeout time.Duration, retryCount int) (*Client, error) {
	cc, err := grpc.NewClient(
		addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to init profile gRPC client: %w", err)
	}

	return &Client{
		api: v1.NewProfileServiceClient(cc),
		log: log,
	}, nil
}
