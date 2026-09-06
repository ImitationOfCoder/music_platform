package app

import (
	"context"
	"os"
	grpc_app "user_microservice/internal/app/grpc"
	profile_repository "user_microservice/internal/repository/profile"
	profile_service "user_microservice/internal/service/profile"
	"user_microservice/pkg/logger"
	"user_microservice/pkg/postgres"
)

type App struct {
	gRPCServer *grpc_app.App
}

func New(ctx context.Context, log *logger.Logger) *App {
	// Postgres
	log.Debug("Initializing postgres connection pool")
	pool, err := postgres.NewConnectionPool(ctx)
	if err != nil {
		log.Error("Failed to init postgres connection pool", log.Err(err))
		os.Exit(1)
	}

	// Repositories
	log.Debug("Initializing repositories")
	profileRepository := profile_repository.NewRepository(pool)

	// Services
	log.Debug("Initializing services")
	profileService := profile_service.NewService(profileRepository)

	gRPCApp, err := grpc_app.New(log.Logger, profileService)
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
