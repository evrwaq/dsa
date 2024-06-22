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

func TestStackPush(t *testing.T) {
	stack := data_structures.NewStack()

	t.Run("Push elements", func(t *testing.T) {
		stack.Push(1)
		if stack.Size() != 1 {
			t.Errorf("expected size 1, got %d", stack.Size())
		}

		stack.Push(2)
		if stack.Size() != 2 {
			t.Errorf("expected size 2, got %d", stack.Size())
		}

		stack.Push(3)
		if stack.Size() != 3 {
			t.Errorf("expected size 3, got %d", stack.Size())
		}

		if stack.IsEmpty() {
			t.Error("expected stack to not be empty")
		}
	})
}
