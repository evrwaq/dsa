package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestNewStack(t *testing.T) {
	stack := data_structures.NewStack()

	t.Run("Initial Size", func(t *testing.T) {
		if stack.Size() != 0 {
			t.Errorf("expected size 0, got %d", stack.Size())
		}
	})

	t.Run("IsEmpty on new stack", func(t *testing.T) {
		if !stack.IsEmpty() {
			t.Error("expected stack to be empty")
		}
	})
}
