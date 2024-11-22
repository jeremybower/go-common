package env

import (
	"errors"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

var ErrEmpty = errors.New("value must not be empty")

func NotEmpty(v string) error {
	if strings.TrimSpace(v) == "" {
		return ErrEmpty
	}

	return nil
}

var ErrNonNegative = errors.New("value must be non-negative")

func NonNegative[T Number](v T) error {
	if v < T(0) {
		return ErrNonNegative
	}

	return nil
}

var ErrPositive = errors.New("value must be positive")

func Positive[T Number](v T) error {
	if v <= T(0) {
		return ErrPositive
	}

	return nil
}
