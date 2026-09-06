package grpc_clients

import (
	"api_gateway_microservice/config"
	account_grpc_client "api_gateway_microservice/internal/client/grpc/account"
	profile_grpc_client "api_gateway_microservice/internal/client/grpc/profile"
	grpc_lib "api_gateway_microservice/internal/lib/grpc"
	"api_gateway_microservice/pkg/logger"
	"os"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/timeout"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Clients struct {
	Account     AccountClient
	Profile     ProfileClient
	CurrentUser CurrentUserClient
}

func New(cfg *config.Config, log *logger.Logger) *Clients {
	opts := []logging.Option{
		logging.WithLogOnEvents(
			logging.StartCall,
			logging.FinishCall,
			logging.PayloadSent,
			logging.PayloadReceived,
		),
		// Add any other option (check functions starting with logging.With).
	}
	dialOptions := []grpc.DialOption{
		grpc.WithUnaryInterceptor(retry.UnaryClientInterceptor(
			retry.WithMax(cfg.Client.Retries),
			retry.WithPerRetryTimeout(cfg.Client.TimeoutPerRetry),
		)),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithChainUnaryInterceptor(
			timeout.UnaryClientInterceptor(cfg.Client.Timeout),
			logging.UnaryClientInterceptor(InterceptorLogger(log.Logger), opts...),
		),
	}

	accountServiceClient, err := account_grpc_client.New(
		log,
		cfg.Microservices.AccountMicroserviceAddr,
		dialOptions,
	)
	if err != nil {
		log.Error("Failed to init account gRPC client.")
		os.Exit(1)
	}

	profileServiceClient, err := profile_grpc_client.New(
		log,
		cfg.Microservices.ProfileMicroserviceAddr,
		dialOptions,
	)
	if err != nil {
		log.Error("Failed to init profile gRPC client.")
		os.Exit(1)
	}

	return &Clients{
		Account: accountServiceClient,
		Profile: profileServiceClient,
	}
}
