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

func TestRequired(t *testing.T) {
	os.Setenv("TEST_ENV", "test")
	assert.Equal(t, "test", Required("TEST_ENV"))

	assert.Panics(t, func() { Required("MISSING_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "test")
	assert.Panics(t, func() {
		Required("TEST_ENV", func(v string) error {
			return expectedErr
		})
	})
}

func TestRequiredBool(t *testing.T) {
	for _, value := range []string{"t", "T", "true", "TRUE", "True", "1"} {
		os.Setenv("TEST_ENV", value)
		assert.True(t, RequiredBool("TEST_ENV"))
	}

	for _, value := range []string{"f", "F", "false", "FALSE", "False", "0"} {
		os.Setenv("TEST_ENV", value)
		assert.False(t, RequiredBool("TEST_ENV"))
	}

	assert.Panics(t, func() { RequiredBool("MISSING_ENV") })

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { RequiredBool("TEST_ENV") })
}

func TestRequiredFloat32(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatFloat(float64(math.MaxFloat32), 'f', -1, 32))
	assert.Equal(t, float32(math.MaxFloat32), RequiredFloat32("TEST_ENV"))

	assert.Panics(t, func() { RequiredFloat32("MISSING_ENV") })

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { RequiredFloat32("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1.0")
	assert.Panics(t, func() {
		RequiredFloat32("TEST_ENV", func(v float32) error {
			return expectedErr
		})
	})
}

func TestRequiredFloat64(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatFloat(float64(math.MaxFloat64), 'f', -1, 64))
	assert.Equal(t, float64(math.MaxFloat64), RequiredFloat64("TEST_ENV"))

	assert.Panics(t, func() { RequiredFloat64("MISSING_ENV") })

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { RequiredFloat64("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1.0")
	assert.Panics(t, func() {
		RequiredFloat64("TEST_ENV", func(v float64) error {
			return expectedErr
		})
	})
}

func TestRequiredInt(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt), 10))
	assert.Equal(t, int(math.MaxInt), RequiredInt("TEST_ENV"))

	assert.Panics(t, func() { RequiredInt("MISSING_ENV") })

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { RequiredInt("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		RequiredInt("TEST_ENV", func(v int) error {
			return expectedErr
		})
	})
}

func TestRequiredInt32(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt32), 10))
	assert.Equal(t, int32(math.MaxInt32), RequiredInt32("TEST_ENV"))

	assert.Panics(t, func() { RequiredInt32("MISSING_ENV") })

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { RequiredInt32("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		RequiredInt32("TEST_ENV", func(v int32) error {
			return expectedErr
		})
	})
}

func TestRequiredInt64(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt64), 10))
	assert.Equal(t, int64(math.MaxInt64), RequiredInt64("TEST_ENV"))

	assert.Panics(t, func() { RequiredInt64("MISSING_ENV") })

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { RequiredInt64("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		RequiredInt64("TEST_ENV", func(v int64) error {
			return expectedErr
		})
	})
}

func TestRequiredURL(t *testing.T) {
	os.Setenv("TEST_ENV", "https://example.com")
	assert.Equal(t, url.URL{Scheme: "https", Host: "example.com"}, RequiredURL("TEST_ENV"))

	assert.Panics(t, func() { RequiredURL("MISSING_ENV") })

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { RequiredURL("TEST_ENV") })
}
