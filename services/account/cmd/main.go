package main

import (
	"account_microservice/internal/app"
	"account_microservice/pkg/logger"
	"context"
	"os/signal"
	"syscall"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	// Logger
	log := logger.NewLoggerMust()
	defer log.Close()

	log.Debug("Starting application!")
	log.Debug("Initializing gRPC server")

	application := app.New(ctx, log)

	go application.MustRun()

	<-ctx.Done()

	application.GracefulStop()

	log.Debug("Application stopped!")
}
