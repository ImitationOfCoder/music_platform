package grpc_app

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Port int `envconfig:"GRPC_PROFILE_SERVICE_PORT"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process gRPC config: %w", err)
	}

	return config, nil
}
