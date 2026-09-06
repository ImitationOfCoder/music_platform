package profile_repository

import (
	"user_microservice/pkg/postgres"
)

type ProfileRepository struct {
	pool postgres.Pool
}

func NewRepository(
	pool postgres.Pool,
) *ProfileRepository {
	return &ProfileRepository{
		pool: pool,
	}
}
