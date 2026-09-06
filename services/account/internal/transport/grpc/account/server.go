package account_grpc

import (
	"account_microservice/internal/domain"
	"account_microservice/pkg/logger"
	"context"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc"
)

type AccountService interface {
	GetAccountById(ctx context.Context, id int64) (*domain.Account, error)
	CreateAccount(ctx context.Context, name string, email string, password string) error
	Login(ctx context.Context, email string, password string) (*domain.Account, error)
}

type ServerApi struct {
	v1.UnimplementedAccountServiceServer
	log            *logger.Logger
	accountService AccountService
}

func Register(gRPC *grpc.Server, log *logger.Logger, accountService AccountService) {
	v1.RegisterAccountServiceServer(gRPC, &ServerApi{
		log:            log,
		accountService: accountService,
	})
}
