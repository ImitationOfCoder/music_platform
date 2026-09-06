package grpc_clients

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Timeout         time.Duration `env:"CLIENT_TIMEOUT" default:"10s"`
	Retries         uint          `env:"CLIENT_RETRIES" default:"3"`
	TimeoutPerRetry time.Duration `env:"CLIENT_TIMEOUT_PER_RETRY" default:"2s"`

	Microservices struct {
		AccountMicroserviceAddr string `env:"ACCOUNT_MICROSERVICE_ADDR,required"`
		ProfileMicroserviceAddr string `env:"PROFILE_MICROSERVICE_ADDR,required"`
	}
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
