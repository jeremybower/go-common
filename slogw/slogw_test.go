package slogw

import (
	"log/slog"
	"regexp"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSlogWriter(t *testing.T) {
	t.Parallel()

	var sb strings.Builder
	handler := slog.NewTextHandler(&sb, nil)
	logger := slog.New(handler)
	writer := New(slog.LevelInfo, logger)

	n, err := writer.Write([]byte("test"))
	require.NoError(t, err)
	require.Equal(t, 4, n)

	re := regexp.MustCompile(`time=(.+) level=(.+) msg=(.+)`)
	matches := re.FindStringSubmatch(sb.String())
	require.Len(t, matches, 4)

	loggedAt, err := time.Parse(time.RFC3339, matches[1])
	require.NoError(t, err)
	assert.WithinDuration(t, time.Now(), loggedAt, 5*time.Second)

	assert.Equal(t, "INFO", matches[2])
	assert.Equal(t, "test", matches[3])
}
