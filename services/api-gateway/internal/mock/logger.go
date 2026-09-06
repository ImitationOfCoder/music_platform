package mock

import (
	"api_gateway_microservice/pkg/logger"
	"io"
	"log/slog"
)

func NewTestLogger() *logger.Logger {
	return &logger.Logger{
		Logger: slog.New(slog.NewTextHandler(io.Discard, nil)),
	}
}
