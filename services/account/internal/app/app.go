package app

import (
	grpc_app "account_microservice/internal/app/grpc"
	profile_service_grpc_client "account_microservice/internal/clients/profile_service/grpc"
	account_repository "account_microservice/internal/repository/account"
	account_service "account_microservice/internal/service/account"
	"account_microservice/pkg/logger"
	"account_microservice/pkg/postgres"
	"account_microservice/pkg/snowflake"
	"context"
	"os"
	"time"
)

type App struct {
	gRPCServer *grpc_app.App
}

func New(ctx context.Context, log *logger.Logger) *App {
	snow, err := snowflake.NewSnowflakeGenerator()
	if err != nil {
		log.Error("Failed to init snowflake generator", log.Err(err))
		os.Exit(1)
	}

	// Postgres
	log.Debug("Initializing postgres connection pool")
	pool, err := postgres.NewConnectionPool(ctx)
	if err != nil {
		log.Error("Failed to init postgres connection pool", log.Err(err))
		os.Exit(1)
	}

	// Repositories
	log.Debug("Initializing repositories")
	accountRepository := account_repository.NewRepository(pool)

	// Clients
	log.Debug("Initializing client")
	// TODO: Вынести адрес микросервиса в конфиг
	profileServiceClient, err := profile_service_grpc_client.New(ctx, log, "profile-microservice:4000", 3*time.Second, 3)
	if err != nil {
		log.Error("Failed to init UserService gRPC client.")
		os.Exit(1)
	}

	// Services
	log.Debug("Initializing services")
	accountService := account_service.NewService(
		log,
		snow,
		accountRepository,
		profileServiceClient,
	)

	gRPCApp, err := grpc_app.New(log, accountService)
	if err != nil {
		log.Error("Failed to init gRPC server", log.Err(err))
		os.Exit(1)
	}

	return &App{
		gRPCServer: gRPCApp,
	}
}

func (a *App) MustRun() {
	a.gRPCServer.MustRun()
}

func (a *App) GracefulStop() {
	a.gRPCServer.Stop()
}
