package account_grpc

import (
	"account_microservice/internal/domain"
	"context"
	"errors"
	"fmt"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *ServerApi) GetAccountById(ctx context.Context, req *v1.GetAccountByIdRequest) (*v1.GetAccountByIdResponse, error) {
	id := req.GetId()

	account, err := s.accountService.GetAccountById(ctx, id)
	if err != nil {
		if errors.Is(err, domain.ErrAccountNotFound) {
			return nil, status.Error(codes.NotFound, "Account not found.")
		}

		return nil, status.Error(codes.Internal, fmt.Sprintf("Internal Server Error: %s", err.Error()))
	}

	return &v1.GetAccountByIdResponse{
		Id:    account.Id,
		Email: account.Email,
	}, nil
}
