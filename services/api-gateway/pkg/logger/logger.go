package logger

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Logger struct {
	*slog.Logger

	file *os.File
}

type Config struct {
	Env    string `envconfig:"APP_ENV" required:"true"`
	Level  string `envconfig:"LOGGER_LEVEL" required:"true"`
	Folder string `envconfig:"LOGGER_FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return Config{}, fmt.Errorf("process config: %w", err)
	}

	return config, nil
}

func FromContext(ctx context.Context) *Logger {
	log, ok := ctx.Value("log").(*Logger)

	if !ok {
		panic("no logger in context")
	}

	return log
}

func NewLoggerMust() *Logger {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("get logger config: %w", err)
		panic(err)
	}

	log, err := NewLogger(config)

	if err != nil {
		fmt.Println("failed to init application logger:", err)
		os.Exit(1)
	}

	return log
}

func NewLogger(config Config) (*Logger, error) {
	var handler slog.Handler
	var err error

	logFile, err := getLogFile(config.Folder)

	if err != nil {
		return nil, fmt.Errorf("create log file: %w", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)

	switch config.Env {
	case "local":
		handler = getLocalHandler(multiWriter)
	case "dev":
		handler = getDevHandler(multiWriter)
	default:
		handler = getLocalHandler(multiWriter)
	}

	logger := slog.New(handler)
	slog.SetDefault(logger)

	return &Logger{
		Logger: logger,
		file:   logFile,
	}, nil
}

func getLogFile(folder string) (*os.File, error) {
	timestamp := time.Now().UTC().Format("2006-01-02T15-04-05.00000")
	logFilePath := filepath.Join(
		folder,
		fmt.Sprintf("%s.log", timestamp),
	)

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return nil, fmt.Errorf("open log file: %w", err)
	}

	return logFile, nil
}

func getLocalHandler(writer io.Writer) *slog.TextHandler {
	return slog.NewTextHandler(
		writer,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	)
}

func getDevHandler(writer io.Writer) *slog.JSONHandler {
	return slog.NewJSONHandler(
		writer,
		&slog.HandlerOptions{
			Level: slog.LevelDebug,
		},
	)
}

func (l *Logger) With(fields ...slog.Attr) *Logger {
	logger := l.Logger

	for _, field := range fields {
		logger = logger.With(field)
	}

	return &Logger{
		Logger: logger,
		file:   l.file,
	}
}

func (l *Logger) Close() {
	if err := l.file.Close(); err != nil {
		fmt.Println("failed to close application logger:", err)
	}
}

func (l *Logger) Err(err error) slog.Attr {
	return slog.Attr{
		Key:   "error",
		Value: slog.StringValue(err.Error()),
	}
}
