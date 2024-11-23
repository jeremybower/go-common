package env

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

var ErrInvalidPortNumber = errors.New("value must be a valid port number")

func IsPort(v string) error {
	port, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrInvalidPortNumber, err)
	}

	if port < 1 || port > 65535 {
		return fmt.Errorf("%w: %d", ErrInvalidPortNumber, port)
	}

	return nil
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
