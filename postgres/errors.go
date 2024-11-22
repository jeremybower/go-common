package postgres

import (
	"errors"

	"github.com/jackc/pgx/v5"
)

var ErrNotFound = errors.New("not found")

func NormalizeError(err error) error {
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrNotFound
	}

	return err
}
