package app

import (
	"api_gateway_microservice/config"
	grpc_client "api_gateway_microservice/internal/client/grpc"
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	handler "api_gateway_microservice/internal/transport/rest"
	"api_gateway_microservice/pkg/httpserver"
	"api_gateway_microservice/pkg/logger"
	redis_client "api_gateway_microservice/pkg/redis"
	"fmt"
)

type servers struct {
	http *httpserver.Server
}

func initServer(
	cfg *config.Config,
	log *logger.Logger,
	jwtManager *jwt_lib.JWTManager,
	grpcClients *grpc_client.Clients,
) *servers {
	httpServer := httpserver.New(cfg, log)
	handler.NewRouter(httpServer.App, log, jwtManager, grpcClients)

	return &servers{
		http: httpServer,
	}
}

func (s *servers) run() {
	s.http.Start()
	fmt.Println("Hello, world!")
}

func Run(cfg *config.Config) {
	// Logger
	log := logger.NewLoggerMust(cfg)
	defer log.Close()

	// Redis
	rdb := redis_client.NewRedisClient(cfg)
	defer rdb.Close()

	// Clients
	grpcClients := grpc_client.New(cfg, log)

	// jwtManager
	jwtManager := jwt_lib.NewJWTManager(cfg, rdb)

	s := initServer(cfg, log, jwtManager, grpcClients)
	s.run()
}
