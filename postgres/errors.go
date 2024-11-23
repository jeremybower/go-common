package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jeremybower/go-common"
)

func NormalizeError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return common.ErrNotFound
	}

	return err
}
