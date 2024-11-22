package env

import (
	"errors"
	"math"
	"net/url"
	"os"
	"strconv"
	"testing"

	. "github.com/jeremybower/go-common/ext"
	"github.com/stretchr/testify/assert"
)

func TestNilable(t *testing.T) {
	os.Setenv("TEST_ENV", "test")
	v := Nilable("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, "test", *v.Value)

	assert.False(t, Nilable("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[string](), Nilable("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr("test"), Nilable("MISSING_ENV").Or(Ptr("test")))

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "test")
	assert.Panics(t, func() {
		Nilable("TEST_ENV", func(v string) error {
			return expectedErr
		})
	})
}

func TestNilableBool(t *testing.T) {
	for _, value := range []string{"t", "T", "true", "TRUE", "True", "1"} {
		os.Setenv("TEST_ENV", value)
		v := NilableBool("TEST_ENV")
		assert.True(t, v.Valid)
		assert.True(t, *v.Value)
	}

	for _, value := range []string{"f", "F", "false", "FALSE", "False", "0"} {
		os.Setenv("TEST_ENV", value)
		v := NilableBool("TEST_ENV")
		assert.True(t, v.Valid)
		assert.False(t, *v.Value)
	}

	assert.False(t, NilableBool("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[bool](), NilableBool("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr(true), NilableBool("MISSING_ENV").Or(Ptr(true)))

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { NilableBool("TEST_ENV") })
}

func TestNilableFloat32(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatFloat(float64(math.MaxFloat32), 'f', -1, 32))
	v := NilableFloat32("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, float32(math.MaxFloat32), *v.Value)

	assert.False(t, NilableFloat32("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[float32](), NilableFloat32("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr(float32(math.MaxFloat32)), NilableFloat32("MISSING_ENV").Or(Ptr(float32(math.MaxFloat32))))

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { NilableFloat32("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1.0")
	assert.Panics(t, func() {
		NilableFloat32("TEST_ENV", func(v float32) error {
			return expectedErr
		})
	})
}

func TestNilableFloat64(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatFloat(float64(math.MaxFloat64), 'f', -1, 64))
	v := NilableFloat64("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, float64(math.MaxFloat64), *v.Value)

	assert.False(t, NilableFloat64("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[float64](), NilableFloat64("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr(float64(math.MaxFloat64)), NilableFloat64("MISSING_ENV").Or(Ptr(float64(math.MaxFloat64))))

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { NilableFloat64("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1.0")
	assert.Panics(t, func() {
		NilableFloat64("TEST_ENV", func(v float64) error {
			return expectedErr
		})
	})
}

func TestNilableInt(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt), 10))
	v := NilableInt("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, int(math.MaxInt), *v.Value)

	assert.False(t, NilableInt("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[int](), NilableInt("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr(int(math.MaxInt)), NilableInt("MISSING_ENV").Or(Ptr(int(math.MaxInt))))

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { NilableInt("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		NilableInt("TEST_ENV", func(v int) error {
			return expectedErr
		})
	})
}

func TestNilableInt32(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt32), 10))
	v := NilableInt32("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, int32(math.MaxInt32), *v.Value)

	assert.False(t, NilableInt32("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[int32](), NilableInt32("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr(int32(math.MaxInt32)), NilableInt32("MISSING_ENV").Or(Ptr(int32(math.MaxInt32))))

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { NilableInt32("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		NilableInt32("TEST_ENV", func(v int32) error {
			return expectedErr
		})
	})
}

func TestNilableInt64(t *testing.T) {
	os.Setenv("TEST_ENV", strconv.FormatInt(int64(math.MaxInt64), 10))
	v := NilableInt64("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, int64(math.MaxInt64), *v.Value)

	assert.False(t, NilableInt64("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[int64](), NilableInt64("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr(int64(math.MaxInt64)), NilableInt64("MISSING_ENV").Or(Ptr(int64(math.MaxInt64))))

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { NilableInt64("TEST_ENV") })

	expectedErr := errors.New("expected")
	os.Setenv("TEST_ENV", "1")
	assert.Panics(t, func() {
		NilableInt64("TEST_ENV", func(v int64) error {
			return expectedErr
		})
	})
}

func TestNilableURL(t *testing.T) {
	u := url.URL{Scheme: "https", Host: "example.com"}
	os.Setenv("TEST_ENV", u.String())
	v := NilableURL("TEST_ENV")
	assert.True(t, v.Valid)
	assert.Equal(t, &u, v.Value)

	assert.False(t, NilableURL("MISSING_ENV").Valid)
	assert.Equal(t, NilPtr[url.URL](), NilableURL("MISSING_ENV").Or(nil))
	assert.Equal(t, Ptr(u), NilableURL("MISSING_ENV").Or(Ptr(u)))

	os.Setenv("TEST_ENV", "invalid")
	assert.Panics(t, func() { NilableURL("TEST_ENV") })
}
