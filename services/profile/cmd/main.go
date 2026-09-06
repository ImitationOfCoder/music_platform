package main

import (
	"context"
	"os/signal"
	"syscall"
	"user_microservice/internal/app"
	"user_microservice/pkg/logger"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer cancel()

	// Logger
	log := logger.NewLoggerMust()
	defer log.Close()

	log.Debug("Starting application!")

	application := app.New(ctx, log)

	go application.MustRun()

	<-ctx.Done()

	application.GracefulStop()

	log.Debug("Application stopped!")
}
