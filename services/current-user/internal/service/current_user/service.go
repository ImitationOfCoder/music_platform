package current_user_service

import (
	grpc_clients "curret_user_microservice/internal/client/grpc"
)

type CurrentUserService struct {
	clients *grpc_clients.Clients
}

func New(
	clients *grpc_clients.Clients,
) *CurrentUserService {
	return &CurrentUserService{
		clients: clients,
	}
}
