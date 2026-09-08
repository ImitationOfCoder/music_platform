package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnectionPool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

func New(
	ctx context.Context,
	user string,
	password string,
	host string,
	port int,
	database string,
	opTimeout time.Duration,
) (*ConnectionPool, error) {
	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user,
		password,
		host,
		port,
		database,
	)

	pgxConfig, err := pgxpool.ParseConfig(connectionString)
	if err != nil {
		return nil, fmt.Errorf("parse pgx config: %w", err)
	}

	pgxPool, err := pgxpool.NewWithConfig(ctx, pgxConfig)
	if err != nil {
		return nil, fmt.Errorf("create pgx pool: %w", err)
	}

	if err := pgxPool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("pgx pool ping: %w", err)
	}

	return &ConnectionPool{
		Pool:      pgxPool,
		opTimeout: opTimeout,
	}, nil
}

func (p *ConnectionPool) OpTimeout() time.Duration {
	return p.opTimeout
}
