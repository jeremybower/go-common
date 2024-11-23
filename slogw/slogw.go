package slogw

import (
	"context"
	"log/slog"
)

type Writer struct {
	level  slog.Level
	logger *slog.Logger
}

func New(level slog.Level, logger *slog.Logger) *Writer {
	return &Writer{level: level, logger: logger}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.logger.Log(context.Background(), w.level, string(p))
	return len(p), nil
}
