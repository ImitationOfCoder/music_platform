package main

import (
	auth_handler "auth_microservice/internal/handler/auth"
	session_repository "auth_microservice/internal/repositories/session"
	auth_service "auth_microservice/internal/service/auth"
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ImitationOfCoder/music_platform/pkg/http_server"
	"github.com/ImitationOfCoder/music_platform/pkg/logger"
	"github.com/ImitationOfCoder/music_platform/pkg/postgres"
	"github.com/ImitationOfCoder/music_platform/pkg/snowflake"
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

	// Postgres
	log.Debug("Initializing postgres connection pool")
	pool, err := postgres.NewConnectionPool(ctx)
	if err != nil {
		log.Error("Failed to init postgres connection pool", log.Err(err))
		os.Exit(1)
	}

	// Snowflake ID Generator
	log.Debug("Initializing snowflake id generator")
	snow, err := snowflake.NewSnowflakeGenerator()
	if err != nil {
		log.Error("Failed to init snowflake id generator", log.Err(err))
		os.Exit(1)
	}

	// Repositories
	log.Debug("Initializing repositories")
	sessionRepository := session_repository.NewRepository(pool)

	// Services
	log.Debug("Initializing services")
	authService := auth_service.NewService(_, sessionRepository)

	// Handlers
	log.Debug("Initializing handlers")
	authHandler := auth_handler.NewHandler(authService)

	log.Debug("Initializing HTTP server")
	httpServer := http_server.NewHTTPServer(log)

	httpServer.RegisterMiddleware(http_server.RequestID())
	httpServer.RegisterMiddleware(http_server.Logger(log))
	httpServer.RegisterMiddleware(http_server.Panic())
	httpServer.RegisterMiddleware(http_server.Trace())

	var routes []http_server.Route

	routes = append(routes, authHandler.Routes()...)

	apiRouter := http_server.NewAPIRouter(routes...)

	httpServer.RegisterAPIRouters(apiRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("HTTP server run error", log.Err(err))
	}
}
