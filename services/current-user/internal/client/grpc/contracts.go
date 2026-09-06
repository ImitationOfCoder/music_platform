package grpc_clients

import (
	"context"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
)

type (
	AccountClient interface {
		GetAccountById(ctx context.Context, accountId int64) (*v1.GetAccountByIdResponse, error)
		CreateAccount(ctx context.Context, name string, email string, password string) error
		Login(ctx context.Context, email string, password string) (*v1.LoginResponse, error)
	}

	ProfileClient interface {
		GetProfileById(ctx context.Context, id int64) (*v1.GetProfileByIdResponse, error)
	}

	CurrentUserClient interface {
		GetCurrentUser(ctx context.Context) (*v1.GetCurrentUserProfileResponse, error)
	}
)
