package env

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/jeremybower/go-common/optional"
)

func Optional(name string, checks ...func(string) error) optional.Value[string] {
	v, ok := os.LookupEnv(name)
	if !ok {
		return optional.InvalidValue[string]()
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid value for environment variable: %s (%w)", name, err))
		}
	}

	return optional.NewValue(v)
}

func OptionalBool(name string) optional.Value[bool] {
	str := Optional(name, NotEmpty)
	if !str.Valid {
		return optional.InvalidValue[bool]()
	}

	v, err := strconv.ParseBool(str.Value)
	if err != nil {
		panic(fmt.Errorf("invalid boolean value for environment variable: %s (%w)", name, err))
	}

	return optional.NewValue(v)
}

func OptionalFloat32(name string, checks ...func(float32) error) optional.Value[float32] {
	str := Optional(name, NotEmpty)
	if !str.Valid {
		return optional.InvalidValue[float32]()
	}

	v64, err := strconv.ParseFloat(str.Value, 32)
	if err != nil {
		panic(fmt.Errorf("invalid float32 value for environment variable: %s (%w)", name, err))
	}

	v := float32(v64)
	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid float32 value for environment variable: %s (%w)", name, err))
		}
	}

	return optional.NewValue(v)
}

func OptionalFloat64(name string, checks ...func(float64) error) optional.Value[float64] {
	str := Optional(name, NotEmpty)
	if !str.Valid {
		return optional.InvalidValue[float64]()
	}

	v64, err := strconv.ParseFloat(str.Value, 64)
	if err != nil {
		panic(fmt.Errorf("invalid float64 value for environment variable: %s (%w)", name, err))
	}

	for _, check := range checks {
		if err := check(v64); err != nil {
			panic(fmt.Errorf("invalid float64 value for environment variable: %s (%w)", name, err))
		}
	}

	return optional.NewValue(v64)
}

func OptionalInt(name string, checks ...func(int) error) optional.Value[int] {
	str := Optional(name, NotEmpty)
	if !str.Valid {
		return optional.InvalidValue[int]()
	}

	v, err := strconv.Atoi(str.Value)
	if err != nil {
		panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
		}
	}

	return optional.NewValue(v)
}

func OptionalInt32(name string, checks ...func(int32) error) optional.Value[int32] {
	str := Optional(name, NotEmpty)
	if !str.Valid {
		return optional.InvalidValue[int32]()
	}

	v64, err := strconv.ParseInt(str.Value, 10, 32)
	if err != nil {
		panic(fmt.Errorf("invalid int32 value for environment variable: %s (%w)", name, err))
	}

	v := int32(v64)
	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid int32 value for environment variable: %s (%w)", name, err))
		}
	}

	return optional.NewValue(v)
}

func OptionalInt64(name string, checks ...func(int64) error) optional.Value[int64] {
	str := Optional(name, NotEmpty)
	if !str.Valid {
		return optional.InvalidValue[int64]()
	}

	v, err := strconv.ParseInt(str.Value, 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid int64 value for environment variable: %s (%w)", name, err))
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid int64 value for environment variable: %s (%w)", name, err))
		}
	}

	return optional.NewValue(v)
}

func OptionalURL(name string) optional.Value[url.URL] {
	str := Optional(name, NotEmpty)
	if !str.Valid {
		return optional.InvalidValue[url.URL]()
	}

	u, err := url.ParseRequestURI(str.Value)
	if err != nil {
		panic(fmt.Errorf("invalid URL value for environment variable: %s (%w)", name, err))
	}

	return optional.NewValue(*u)
}
