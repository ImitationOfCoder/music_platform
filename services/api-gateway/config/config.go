package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App           app
		HTTP          http
		Logger        logger
		Redis         redis
		Client        client
		JWT           jwt
		Microservices microservices
	}

	app struct {
		Env string `env:"APP_ENV,required"`
	}

	http struct {
		Address string `env:"HTTP_ADDRESS"`
	}

	logger struct {
		Level  string `env:"LOGGER_LEVEL" envDefault:"DEBUG"`
		Folder string `env:"LOGGER_FOLDER" envDefault:"./out/logs"`
	}

	redis struct {
		Host     string `env:"REDIS_HOST" default:"localhost"`
		Port     int    `env:"REDIS_PORT" default:"6379"`
		Password string `env:"REDIS_PASSWORD" default:""`
		Db       int    `env:"REDIS_DB" default:"0"`
	}

	client struct {
		Timeout         time.Duration `env:"CLIENT_TIMEOUT" default:"10s"`
		Retries         uint          `env:"CLIENT_RETRIES" default:"3"`
		TimeoutPerRetry time.Duration `env:"CLIENT_TIMEOUT_PER_RETRY" default:"2s"`
	}

	jwt struct {
		Secret string `env:"JWT_SECRET" default:""`
	}

	microservices struct {
		AccountMicroserviceAddr string `env:"ACCOUNT_MICROSERVICE_ADDR,required"`
		ProfileMicroserviceAddr string `env:"PROFILE_MICROSERVICE_ADDR,required"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}

	return cfg, nil
}
