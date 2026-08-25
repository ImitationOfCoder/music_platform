package user_repository

import (
	"user_microservice/pkg/postgres"
)

type UserRepository struct {
	pool postgres.Pool
}

func NewRepository(
	pool postgres.Pool,
) *UserRepository {
	return &UserRepository{
		pool: pool,
	}
}
