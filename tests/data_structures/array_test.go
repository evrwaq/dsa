package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestArray(t *testing.T) {
	size := 10
	arr := data_structures.NewArray(size)

	t.Run("Initial Size", func(t *testing.T) {
		if arr.Size() != size {
			t.Errorf("expected size %d, got %d", size, arr.Size())
		}
	})

	t.Run("Set and Get First Element", func(t *testing.T) {
		err := arr.Set(0, "test")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		value, err := arr.Get(0)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if value != "test" {
			t.Errorf("expected value %s, got %v", "test", value)
		}
	})

	t.Run("Set Out of Bounds", func(t *testing.T) {
		err := arr.Set(size, "out of bounds")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})

	t.Run("Get Out of Bounds", func(t *testing.T) {
		_, err := arr.Get(size)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
