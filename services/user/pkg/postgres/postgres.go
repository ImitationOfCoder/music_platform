package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kelseyhightower/envconfig"
)

type Pool interface {
	Query(ctx context.Context, sql string, args ...any) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...any) pgx.Row
	Exec(ctx context.Context, sql string, arguments ...any) (pgconn.CommandTag, error)
	Close()

	OpTimeout() time.Duration
}

type ConnectionPool struct {
	*pgxpool.Pool
	opTimeout time.Duration
}

type Config struct {
	Host     string        `envconfig:"POSTGRES_HOST"     required:"true"`
	Port     string        `envconfig:"POSTGRES_PORT"     default:"5432"`
	User     string        `envconfig:"POSTGRES_USER"     required:"true"`
	Password string        `envconfig:"POSTGRES_PASSWORD" required:"true"`
	Database string        `envconfig:"POSTGRES_DATABASE" required:"true"`
	Timeout  time.Duration `envconfig:"POSTGRES_TIMEOUT"  required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process config: %w", err)
	}

	return config, nil
}

func NewConnectionPool(ctx context.Context) (*ConnectionPool, error) {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get postgres config: %w", err)
		panic(err)
	}

	connectionString := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Database,
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
		opTimeout: config.Timeout,
	}, nil
}

func (p *ConnectionPool) OpTimeout() time.Duration {
	return p.opTimeout
}
