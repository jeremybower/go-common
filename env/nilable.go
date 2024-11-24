package env

import (
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/jeremybower/go-common/nilable"
)

func Nilable(name string, checks ...func(string) error) nilable.Value[string] {
	v, ok := os.LookupEnv(name)
	if !ok {
		return nilable.InvalidValue[string]()
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid value for environment variable: %s (%w)", name, err))
		}
	}

	return nilable.NewValue(&v)
}

func NilableBool(name string) nilable.Value[bool] {
	str := Nilable(name, NotEmpty)
	if !str.Valid {
		return nilable.InvalidValue[bool]()
	}

	v, err := strconv.ParseBool(*str.Value)
	if err != nil {
		panic(fmt.Errorf("invalid boolean value for environment variable: %s (%w)", name, err))
	}

	return nilable.NewValue(&v)
}

func NilableFloat32(name string, checks ...func(float32) error) nilable.Value[float32] {
	str := Nilable(name, NotEmpty)
	if !str.Valid {
		return nilable.InvalidValue[float32]()
	}

	v64, err := strconv.ParseFloat(*str.Value, 32)
	if err != nil {
		panic(fmt.Errorf("invalid float32 value for environment variable: %s (%w)", name, err))
	}

	v := float32(v64)
	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid float32 value for environment variable: %s (%w)", name, err))
		}
	}

	return nilable.NewValue(&v)
}

func NilableFloat64(name string, checks ...func(float64) error) nilable.Value[float64] {
	str := Nilable(name, NotEmpty)
	if !str.Valid {
		return nilable.InvalidValue[float64]()
	}

	v64, err := strconv.ParseFloat(*str.Value, 64)
	if err != nil {
		panic(fmt.Errorf("invalid float64 value for environment variable: %s (%w)", name, err))
	}

	for _, check := range checks {
		if err := check(v64); err != nil {
			panic(fmt.Errorf("invalid float64 value for environment variable: %s (%w)", name, err))
		}
	}

	return nilable.NewValue(&v64)
}

func NilableInt(name string, checks ...func(int) error) nilable.Value[int] {
	str := Nilable(name, NotEmpty)
	if !str.Valid {
		return nilable.InvalidValue[int]()
	}

	v, err := strconv.Atoi(*str.Value)
	if err != nil {
		panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
		}
	}

	return nilable.NewValue(&v)
}

func NilableInt32(name string, checks ...func(int32) error) nilable.Value[int32] {
	str := Nilable(name, NotEmpty)
	if !str.Valid {
		return nilable.InvalidValue[int32]()
	}

	v64, err := strconv.ParseInt(*str.Value, 10, 32)
	if err != nil {
		panic(fmt.Errorf("invalid int32 value for environment variable: %s (%w)", name, err))
	}

	v := int32(v64)
	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid int32 value for environment variable: %s (%w)", name, err))
		}
	}

	return nilable.NewValue(&v)
}

func NilableInt64(name string, checks ...func(int64) error) nilable.Value[int64] {
	str := Nilable(name, NotEmpty)
	if !str.Valid {
		return nilable.InvalidValue[int64]()
	}

	v, err := strconv.ParseInt(*str.Value, 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid int64 value for environment variable: %s (%w)", name, err))
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid int64 value for environment variable: %s (%w)", name, err))
		}
	}

	return nilable.NewValue(&v)
}

func NilableURL(name string) nilable.Value[url.URL] {
	str := Nilable(name, NotEmpty)
	if !str.Valid {
		return nilable.InvalidValue[url.URL]()
	}

	u, err := url.ParseRequestURI(*str.Value)
	if err != nil {
		panic(fmt.Errorf("invalid URL value for environment variable: %s (%w)", name, err))
	}

	return nilable.NewValue(u)
}
