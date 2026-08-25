package session_repository

import (
	"github.com/ImitationOfCoder/music_platform/pkg/postgres"
)

type SessionRepository struct {
	pool postgres.Pool
}

func NewRepository(
	pool postgres.Pool,
) *SessionRepository {
	return &SessionRepository{
		pool: pool,
	}
}
