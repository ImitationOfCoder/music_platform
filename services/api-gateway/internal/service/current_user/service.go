package current_user_service

import (
	"api_gateway_microservice/internal/client"
	"api_gateway_microservice/pkg/logger"
)

type Service struct {
	log         *logger.Logger
	gRPCClients *client.GrpcClients
}

func New(log *logger.Logger, gRPCClients *client.GrpcClients) *Service {
	return &Service{
		log:         log,
		gRPCClients: gRPCClients,
	}
}
