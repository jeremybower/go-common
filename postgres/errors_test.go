package postgres

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jeremybower/go-common"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeError(t *testing.T) {
	t.Parallel()

	assert.Equal(t, common.ErrNotFound, NormalizeError(pgx.ErrNoRows))
	assert.Equal(t, assert.AnError, NormalizeError(assert.AnError))
}
