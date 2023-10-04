// Package result provides Result implementation.
//
// A result can be present in two conditions: real and fake.
//
// A real result contains a value and no error, and can be
// safely unwrapped via (*Result).Value() method to retrieve
// the actual value.
//
// A fake result contains an error and no value (or a "nil" value)
// and cannot be unwrapped (unwrapping will lead to a runtime panic).
//
// There is only one way to distinguish a fake result from a real one:
// check for an error present in the result via calling (*Result).Error(),
// and if the method returns a non-nil error, then the result is a fake
// and must not be treated as a real value.
package result

import (
	"fmt"
)

// Result is a result value, that can either contain an
// actual value, or an error.
//
// The value of a Result can be retrieved via (*Result).Value()
// method. Though, before getting the value, one must check if
// the Result real or a fake (via (*Result).Error()), because
// retreiving value from a fake result will lead to a panic.
//
// To create a Result you should use functions: Fake - to create
// a fake result, and Real - to create a real one.
//
// Creating a result directly via struct construction leads to
// undefined behaviour.
type Result[T any] struct {
	value T
	err   error
}

// Fake creates a new "fake" Result.
func Fake[T any](err error) Result[T] {
	if err == nil {
		panic("cannot create a fake result without an error (err is nil)")
	}
	return Result[T]{
		err: err,
	}
}

// Real creates a new "real" Result with a value.
func Real[T any](value T) Result[T] {
	return Result[T]{
		value: value,
	}
}

// Of wraps (T, error) return result into a Result.
func Of[T any](value T, err error) Result[T] {
	if err != nil {
		return Fake[T](err)
	}
	return Real(value)
}

// Error returns the error of the Result, or nil if the Result is "real".
func (r *Result[_]) Error() error {
	return r.err
}

// Value returns the value of the Rsult, or panics with the result is "fake".
func (r *Result[T]) Value() T {
	if r.err != nil {
		panic(fmt.Errorf("the result is fake: %w", r.err))
	}
	return r.value
}
