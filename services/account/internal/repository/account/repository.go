package account_repository

import (
	"account_microservice/pkg/postgres"
)

type AccountRepository struct {
	pool postgres.Pool
}

func NewRepository(
	pool postgres.Pool,
) *AccountRepository {
	return &AccountRepository{
		pool: pool,
	}
}
