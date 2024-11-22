package env

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
)

func Required(name string, checks ...func(string) error) string {
	v, ok := os.LookupEnv(name)
	if !ok {
		panic(fmt.Errorf("missing required environment variable: %s", name))
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid value for environment variable: %s (%w)", name, err))
		}
	}

	return v
}

func RequiredBool(name string) bool {
	str := Required(name, NotEmpty)

	v, err := strconv.ParseBool(str)
	if err != nil {
		panic(fmt.Errorf("invalid boolean value for environment variable: %s (%w)", name, err))
	}

	return v
}

func RequiredFloat32(name string, checks ...func(float32) error) float32 {
	str := Required(name, NotEmpty)

	v64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		panic(fmt.Errorf("invalid float32 value for environment variable: %s (%w)", name, err))
	}

	v := float32(v64)
	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid float32 value for environment variable: %s (%w)", name, err))
		}
	}

	return v
}

func RequiredFloat64(name string, checks ...func(float64) error) float64 {
	str := Required(name, NotEmpty)

	v, err := strconv.ParseFloat(str, 64)
	if err != nil {
		panic(fmt.Errorf("invalid float64 value for environment variable: %s (%w)", name, err))
	}

	for _, check := range checks {
		if err := check(v); err != nil {
			panic(fmt.Errorf("invalid float64 value for environment variable: %s (%w)", name, err))
		}
	}

	return v
}

func RequiredInt(name string, fns ...func(int) error) int {
	str := Required(name, NotEmpty)

	v, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
	}

	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
		}
	}

	return v
}

func RequiredInt32(name string, fns ...func(int32) error) int32 {
	str := Required(name, NotEmpty)

	v64, err := strconv.ParseInt(str, 10, 32)
	if err != nil {
		panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
	}

	v := int32(v64)
	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
		}
	}

	return v
}

func RequiredInt64(name string, fns ...func(int64) error) int64 {
	str := Required(name, NotEmpty)

	v, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
	}

	for _, fn := range fns {
		if err := fn(v); err != nil {
			panic(fmt.Errorf("invalid integer value for environment variable: %s (%w)", name, err))
		}
	}

	return v
}

func RequiredURL(name string) url.URL {
	str := Required(name, NotEmpty)

	v, err := url.ParseRequestURI(str)
	if err != nil {
		panic(fmt.Errorf("invalid URL value for environment variable: %s (%w)", name, err))
	}

	return *v
}
