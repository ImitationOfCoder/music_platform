package httpserver

import (
	"api_gateway_microservice/config"
	"api_gateway_microservice/pkg/logger"
	"context"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v5"
)

type Server struct {
	App             *echo.Echo
	address         string
	shutdownTimeout time.Duration

	log *logger.Logger
}

func New(cfg *config.Config, log *logger.Logger) *Server {
	return &Server{
		App:             echo.New(),
		address:         cfg.HTTP.Address,
		shutdownTimeout: 5 * time.Second,

		log: log,
	}
}

func (s *Server) Start() {
	ctx, stop := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer stop()

	sc := echo.StartConfig{
		Address:         s.address,
		GracefulTimeout: s.shutdownTimeout,
	}
	if err := sc.Start(ctx, s.App); err != nil {
		s.log.Error("failed to start server", "error", err)
	}
}
