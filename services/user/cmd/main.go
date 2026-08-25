package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"user_microservice/config"
	user_handler "user_microservice/internal/handler/user"
	user_repository "user_microservice/internal/repository/user"
	user_service "user_microservice/internal/service/user"
	"user_microservice/pkg/logger"
	"user_microservice/pkg/postgres"
	"user_microservice/pkg/snowflake"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM,
	)
	defer cancel()

	cfg := config.NewConfigMust()

	// Logger
	log := logger.NewLoggerMust()
	defer log.Close()

	log.Debug("Starting application!")

	// Postgres
	pool := InitPostgres(ctx, log)

	// Snowflake ID Generator
	snow := InitSnowflakeIdGenerator(log)

	// Repositories
	log.Debug("Initializing repositories")
	userRepository := user_repository.NewRepository(pool)

	// Services
	log.Debug("Initializing services")
	userService := user_service.NewService(userRepository, snow)

	// Handlers
	log.Debug("Initializing handlers")
	userHandler := user_handler.NewHandler(userService)

	log.Debug("Initializing HTTP server")
	e := echo.New()

	e.Logger = log.Logger

	e.Use(middleware.RequestID())

	e.POST("/api/users", userHandler.CreateUser)
	e.GET("/api/users/:id", userHandler.GetUserByID)

	if err := e.Start(cfg.HttpAddr); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}

func InitPostgres(ctx context.Context, log *logger.Logger) *postgres.ConnectionPool {
	// Postgres
	log.Debug("Initializing postgres connection pool")
	pool, err := postgres.NewConnectionPool(ctx)
	if err != nil {
		log.Error("Failed to init postgres connection pool", log.Err(err))
		os.Exit(1)
	}

	return pool
}

func InitSnowflakeIdGenerator(log *logger.Logger) *snowflake.SnowflakeGenerator {
	log.Debug("Initializing snowflake id generator")
	snow, err := snowflake.NewSnowflakeGenerator()
	if err != nil {
		log.Error("Failed to init snowflake id generator", log.Err(err))
		os.Exit(1)
	}

	return snow
}
