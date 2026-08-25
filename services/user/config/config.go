package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	HttpAddr string `envconfig:"HTTP_ADDR"`
}

func NewConfigMust() Config {
	cfg, err := NewConfig()

	if err != nil {
		panic("process config error")
	}

	return cfg
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process config: %w", err)
	}

	return config, nil
}
