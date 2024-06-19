package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestArray(t *testing.T) {
	size := 10
	arr := data_structures.NewArray(size)

	if arr.Size() != size {
		t.Errorf("expected size %d, got %d", size, arr.Size())
	}

	err := arr.Set(0, "test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
