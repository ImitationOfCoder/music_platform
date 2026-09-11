package account_grpc_controller

import (
	"account_microservice/internal/service"

	"github.com/ImitationOfCoder/music_platform/pkg/logger"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc"
)

type ServerApi struct {
	v1.UnimplementedAccountServiceServer
	log            *logger.Logger
	accountService service.AccountService
}

func Register(gRPC *grpc.Server, log *logger.Logger, accountService service.AccountService) {
	v1.RegisterAccountServiceServer(gRPC, &ServerApi{
		log:            log,
		accountService: accountService,
	})
}
