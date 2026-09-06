package app

import (
	"context"
	grpc_app "curret_user_microservice/internal/app/grpc"
	"curret_user_microservice/pkg/logger"
	"os"
)

type App struct {
	gRPCServer *grpc_app.App
}

func New(ctx context.Context, log *logger.Logger) *App {
	gRPCApp, err := grpc_app.New(log)
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
