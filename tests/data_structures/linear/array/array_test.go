package data_structures_linear_test

import (
	ds_errors "dsa/src/data_structures/errors"
	ds "dsa/src/data_structures/linear/array"
	"testing"
)

// TestNewArray tests the creation of a new array and its initial size.
func TestNewArray(t *testing.T) {
	size := 10                 // Define the size of the array.
	array := ds.NewArray(size) // Create a new array with the specified size.

	// Test the initial size of the array.
	t.Run("Initial Size", func(t *testing.T) {
		// Check if the size of the array is as expected.
		if array.Size() != size {
			t.Errorf("expected size %d, got %d", size, array.Size())
		}
	})
}

// TestArraySetAndGet tests setting and getting values in the array.
func TestArraySetAndGet(t *testing.T) {
	size := 10                 // Define the size of the array.
	array := ds.NewArray(size) // Create a new array with the specified size.

	// Test setting and getting the first element in the array.
	t.Run("Set and Get First Element", func(t *testing.T) {
		// Set the value at the first index.
		error := array.Set(0, "test")
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}

		// Get the value at the first index.
		value, error := array.Get(0)
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}

		// Check if the value is as expected.
		if value != "test" {
			t.Errorf("expected value %s, got %v", "test", value)
		}
	})

	// Test setting a value at an out-of-bounds index.
	t.Run("Set Out of Bounds", func(t *testing.T) {
		// Try to set a value at an index equal to the size of the array, which is out of bounds.
		error := array.Set(size, "out of bounds")
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})

	// Test getting a value at an out-of-bounds index.
	t.Run("Get Out of Bounds", func(t *testing.T) {
		// Try to get a value at an index equal to the size of the array, which is out of bounds.
		_, error := array.Get(size)
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}
