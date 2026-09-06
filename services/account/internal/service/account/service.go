package account_service

import (
	"context"

	"account_microservice/internal/domain"
	"account_microservice/pkg/logger"
	"account_microservice/pkg/snowflake"

	v1 "github.com/ImitationOfCoder/music_platform_proto/go"
)

type AccountService struct {
	log                  *logger.Logger
	snowflake            *snowflake.SnowflakeGenerator
	accountRepository    AccountRepository
	userRepository       UserRepository
	profileServiceClient ProfileServiceClient
}

type AccountRepository interface {
	GetAccountById(ctx context.Context, id int64) (*domain.Account, error)
	CreateAccount(ctx context.Context, accountId int64, email string, passwordHash string) error
	GetAccountByEmail(ctx context.Context, email string) (*domain.Account, error)
}

type UserRepository interface {
	GetUserBySessionToken(
		ctx context.Context,
		sessionToken string,
	) (domain.User, error)
}

type ProfileServiceClient interface {
	CreateProfile(ctx context.Context, accountId int64, name string) (*v1.CreateProfileResponse, error)
}

func NewService(
	logger *logger.Logger,
	snowflake *snowflake.SnowflakeGenerator,
	accountRepository AccountRepository,
	profileServiceClient ProfileServiceClient,
) *AccountService {
	return &AccountService{
		log:                  logger,
		snowflake:            snowflake,
		accountRepository:    accountRepository,
		profileServiceClient: profileServiceClient,
	}
}
