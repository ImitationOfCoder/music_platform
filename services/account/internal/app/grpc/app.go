package grpc_app

import (
	account_grpc "account_microservice/internal/transport/grpc/account"
	"account_microservice/pkg/logger"
	"fmt"
	"log/slog"
	"net"

	"google.golang.org/grpc"
)

type App struct {
	log        *logger.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *logger.Logger, accountService account_grpc.AccountService) (*App, error) {
	config, err := NewConfig()
	if err != nil {
		return nil, fmt.Errorf("gRPC config: %w", err)
	}

	gRPCServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			loggingInterceptor,
		),
	)

	account_grpc.Register(gRPCServer, log, accountService)

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
