package auth_service

import (
	"api_gateway_microservice/internal/client"
	jwt_lib "api_gateway_microservice/internal/lib/jwt"
	"api_gateway_microservice/pkg/logger"
)

type Service struct {
	log         *logger.Logger
	jwtManager  *jwt_lib.JWTManager
	gRPCClients *client.GrpcClients
}

func New(
	log *logger.Logger,
	jwtManager *jwt_lib.JWTManager,
	gRPCClients *client.GrpcClients,
) *Service {
	return &Service{
		log:         log,
		jwtManager:  jwtManager,
		gRPCClients: gRPCClients,
	}
}
