package grpc_transport

import (
	"fmt"
	"log/slog"
	"net"
	profile_grpc "profile_microservice/internal/controller/grpc/profile"

	"github.com/ImitationOfCoder/music_platform/pkg/logger"
	"github.com/ImitationOfCoder/music_platform/pkg/postgres"
	"google.golang.org/grpc"
)

type Server struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(port int, log *logger.Logger, pool postgres.Pool) (*Server, error) {
	gRPCServer := grpc.NewServer()

	profile_grpc.Register(gRPCServer, log, pool)

	return &Server{
		log:        log.Logger.With(slog.Int("port", port)),
		gRPCServer: gRPCServer,
		port:       port,
	}, nil
}

func (s *Server) MustRun() {
	if err := s.Run(); err != nil {
		panic(err)
	}
}

func (s *Server) Run() error {
	s.log.Info("Starting gRPC server.")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	if err := s.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("s.gRPCServer.Serve: %w", err)
	}

	return nil
}

func (s *Server) Stop() {
	s.log.Info("Stopping gRPC server.")

	s.gRPCServer.GracefulStop()
}
