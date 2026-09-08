package logger

import (
	"fmt"
	"os"
)

type Env string

func (e Env) String() string {
	return string(e)
}

func (e Env) IsValid() bool {
	switch e {
	case Dev, Prod, Local:
		return true
	default:
		return false
	}
}

const (
	Local Env = "local"
	Dev   Env = "dev"
	Prod  Env = "prod"
)

type Config struct {
	Env    Env
	Folder string
}

func newConfig(env Env, folder string) (*Config, error) {
	if !env.IsValid() {
		return nil, fmt.Errorf("env is invalid")
	}

	if !isValidDirectory(folder) {
		return nil, fmt.Errorf("folder path is invalid")
	}

	return &Config{
		Env:    env,
		Folder: folder,
	}, nil
}

func isValidDirectory(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}

	return fileInfo.IsDir()
}
