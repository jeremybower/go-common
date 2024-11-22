package env

import (
	"errors"
	"math"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOptional(t *testing.T) {
	os.Setenv("TEST_ENV", "test")
	v := Optional("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, "test", v.Value)

	assert.False(t, Optional("MISSING_ENV").Valid)
	assert.Equal(t, "test", Optional("MISSING_ENV").Or("test"))

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "test")
	assert.Panics(t, func() {
		Optional("TEST_ENV", func(v string) error {
			return expectedErr
		})
	})
}

func TestOptionalBool(t *testing.T) {
	for _, value := range []string{"t", "T", "true", "TRUE", "True", "1"} {
		os.Setenv("TEST_ENV", value)
		v := OptionalBool("TEST_ENV")
		assert.True(t, v.Valid)
		assert.True(t, v.Value)
	}

	for _, value := range []string{"f", "F", "false", "FALSE", "False", "0"} {
		os.Setenv("TEST_ENV", value)
		v := OptionalBool("TEST_ENV")
		assert.True(t, v.Valid)
		assert.False(t, v.Value)
	}

	assert.False(t, OptionalBool("MISSING_ENV").Valid)

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { OptionalBool("TEST_ENV") })
}

func TestOptionalFloat32(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatFloat(float64(math.MaxFloat32), 'f', -1, 32))
	v := OptionalFloat32("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, float32(math.MaxFloat32), v.Value)

	assert.False(t, OptionalFloat32("MISSING_ENV").Valid)

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { OptionalFloat32("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1.0")
	assert.Panics(t, func() {
		OptionalFloat32("TEST_ENV", func(v float32) error {
			return expectedErr
		})
	})
}

func TestOptionalFloat64(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatFloat(float64(math.MaxFloat64), 'f', -1, 64))
	v := OptionalFloat64("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, float64(math.MaxFloat64), v.Value)

	assert.False(t, OptionalFloat64("MISSING_ENV").Valid)

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { OptionalFloat64("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1.0")
	assert.Panics(t, func() {
		OptionalFloat64("TEST_ENV", func(v float64) error {
			return expectedErr
		})
	})
}

func TestOptionalInt(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt), 10))
	v := OptionalInt("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, int(math.MaxInt), v.Value)

	assert.False(t, OptionalInt("MISSING_ENV").Valid)

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { OptionalInt("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		OptionalInt("TEST_ENV", func(v int) error {
			return expectedErr
		})
	})
}

func TestOptionalInt32(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt32), 10))
	v := OptionalInt32("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, int32(math.MaxInt32), v.Value)

	assert.False(t, OptionalInt32("MISSING_ENV").Valid)

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { OptionalInt32("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		OptionalInt32("TEST_ENV", func(v int32) error {
			return expectedErr
		})
	})
}

func TestOptionalInt64(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt64), 10))
	v := OptionalInt64("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, int64(math.MaxInt64), v.Value)

	assert.False(t, OptionalInt64("MISSING_ENV").Valid)

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { OptionalInt64("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		OptionalInt64("TEST_ENV", func(v int64) error {
			return expectedErr
		})
	})
}

func TestOptionalURL(t *testing.T) {
	os.Setenv("TEST_ENV", "https://example.com")
	v := OptionalURL("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, url.URL{Scheme: "https", Host: "example.com"}, v.Value)

	assert.False(t, OptionalURL("MISSING_ENV").Valid)

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { OptionalURL("TEST_ENV") })
}
