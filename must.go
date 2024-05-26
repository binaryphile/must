package must

import (
	"fmt"
	"os"
)

// AssertNil panics if err is not nil.
func AssertNil(err error) {
	if err != nil {
		panic(err)
	}
}

// Must returns the value of a (value, error) pair of arguments unless error is non-nil.
// In that case, it panics.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}

	return t
}

// MustGetenv returns the value in the environment variable named by key.
// It panics if the environment variable doesn't exist or is empty.
func MustGetenv(key string) string {
	result := os.Getenv(key)

	if result == "" {
		panic(fmt.Sprint("expected value for environment variable ", key))
	}

	return result
}

func Must1[T, R any](fn func(T) (R, error)) func(T) R {
	return func(t T) R {
		r, err := fn(t)
		if err != nil {
			panic(err)
		}

		return r
	}
}

func Must2[T, T2, R any](fn func(T, T2) (R, error)) func(T, T2) R {
	return func(t T, t2 T2) R {
		r, err := fn(t, t2)
		if err != nil {
			panic(err)
		}

		return r
	}
}
