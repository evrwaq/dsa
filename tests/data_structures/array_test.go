package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestArray(t *testing.T) {
	size := 10
	array := data_structures.NewArray(size)

	t.Run("Initial Size", func(t *testing.T) {
		if array.Size() != size {
			t.Errorf("expected size %d, got %d", size, array.Size())
		}
	})

	t.Run("Set and Get First Element", func(t *testing.T) {
		error := array.Set(0, "test")
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}

		value, error := array.Get(0)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}

		if value != "test" {
			t.Errorf("expected value %s, got %v", "test", value)
		}
	})

	t.Run("Set Out of Bounds", func(t *testing.T) {
		error := array.Set(size, "out of bounds")
		if error == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("Get Out of Bounds", func(t *testing.T) {
		_, error := array.Get(size)
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}
