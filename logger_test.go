package common

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogger(t *testing.T) {
	t.Parallel()

	ctx := context.Background()
	_, err := Logger(ctx)
	assert.ErrorIs(t, err, ErrLoggerNotSet)

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	ctx = WithLogger(ctx, logger)

	logger, err = Logger(ctx)
	assert.NoError(t, err)
	assert.Equal(t, logger, logger)
}
