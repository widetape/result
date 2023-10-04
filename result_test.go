package result_test

import (
	"errors"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/widetape/result"
)

func TestResult(t *testing.T) {
	t.Run("Real", func(t *testing.T) {
		t.Run("returns correct value", func(t *testing.T) {
			value := rand.Int()
			r := result.Real(value)
			assert.Equal(t, value, r.Value(), "expected real result to returns the actual value")
		})
		t.Run("returns no error", func(t *testing.T) {
			r := result.Real[any](nil)
			assert.Nil(t, r.Error())
		})
	})
	t.Run("Fake", func(t *testing.T) {
		t.Run("panics without error", func(t *testing.T) {
			assert.Panics(t, func() { result.Fake[any](nil) }, "expected Fake to panic without error on construction")
		})
		t.Run("returns correct error", func(t *testing.T) {
			err := errors.New("test")
			r := result.Fake[any](err)
			assert.ErrorIs(t, err, r.Error(), "expected to return correct error")
		})
		t.Run("panics on unwrapping", func(t *testing.T) {
			r := result.Fake[any](errors.New("test"))
			assert.Panics(t, func() { r.Value() }, "expected fake result to panic")
		})
	})
	t.Run("Of", func(t *testing.T) {
		alwaysErr := func() (int, error) {
			return 0, errors.New("err")
		}
		neverErr := func() (int, error) {
			return 50, nil
		}
		t.Run("wraps error correctly", func(t *testing.T) {
			r := result.Of(alwaysErr())
			assert.NotNil(t, r.Error(), "expected to get an error")
			assert.Panics(
				t,
				func() {
					r.Value()
				},
			)
		})
		t.Run("wraps value correctly", func(t *testing.T) {
			r := result.Of(neverErr())
			assert.Nil(t, r.Error(), "did not expect an error")
			assert.Equal(t, 50, r.Value(), "the value is not correct")
		})
	})
}
