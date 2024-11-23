package common

import (
	"context"
	"errors"
	"log/slog"
)

type contextKey string

const loggerKey = contextKey("logger")

var ErrLoggerNotSet = errors.New("logger not set on context")

func WithLogger(ctx context.Context, logger *slog.Logger) context.Context {
	return context.WithValue(ctx, loggerKey, logger)
}

func Logger(ctx context.Context) (*slog.Logger, error) {
	logger, ok := ctx.Value(loggerKey).(*slog.Logger)
	if !ok {
		return nil, ErrLoggerNotSet
	}

	return logger, nil
}
