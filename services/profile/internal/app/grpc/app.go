package grpc_app

import (
	"fmt"
	"log/slog"
	"net"
	profile_grpc "user_microservice/internal/transport/grpc/profile"

	"google.golang.org/grpc"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, profileService profile_grpc.ProfileService) (*App, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, fmt.Errorf("gRPC config: %w", err)
	}

	gRPCServer := grpc.NewServer()

	profile_grpc.Register(gRPCServer, profileService)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       config.Port,
	}, nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {
	const op = "grpc_app.Run"

	log := a.log.With(slog.String("op", op), slog.Int("port", a.port))

	log.Info("Starting gRPC server.")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (a *App) Stop() {
	const op = "grpc_app.Stop"

	a.log.With(slog.String("op", op)).Info("Stopping gRPC server.", slog.Int("port", a.port))

	a.gRPCServer.GracefulStop()
}
