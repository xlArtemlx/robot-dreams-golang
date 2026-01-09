package logging

import (
	"io"
	"log/slog"
)

type Config struct {
	Service     string
	Environment string
	Version     string
	Level       slog.Level
	Output      io.Writer
}

func New(cfg Config) *slog.Logger {
	if cfg.Output == nil {
		cfg.Output = io.Discard
	}

	handler := slog.NewJSONHandler(cfg.Output, &slog.HandlerOptions{
		Level:     cfg.Level,
		AddSource: true,
	})

	return slog.New(handler).With(
		slog.String("service", cfg.Service),
		slog.String("environment", cfg.Environment),
		slog.String("version", cfg.Version),
	)
}
