package data_structures_linear_test

import (
	ds_errors "dsa/src/data_structures/errors"
	ds "dsa/src/data_structures/linear/array"
	"testing"
)

func TestNewArray(t *testing.T) {
	size := 10
	array := ds.NewArray(size)

	t.Run("Initial Size", func(t *testing.T) {
		if array.Size() != size {
			t.Errorf("expected size %d, got %d", size, array.Size())
		}
	})
}

func TestArraySetAndGet(t *testing.T) {
	size := 10
	array := ds.NewArray(size)

	t.Run("Set and Get First Element", func(t *testing.T) {
		error := array.Set(0, "test")
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}

		value, error := array.Get(0)
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}

		if value != "test" {
			t.Errorf("expected value %s, got %v", "test", value)
		}
	})

	t.Run("Set Out of Bounds", func(t *testing.T) {
		error := array.Set(size, "out of bounds")
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})

	t.Run("Get Out of Bounds", func(t *testing.T) {
		_, error := array.Get(size)
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}
