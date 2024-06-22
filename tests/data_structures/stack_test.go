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

func TestStackPop(t *testing.T) {
	stack := data_structures.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Run("Pop elements", func(t *testing.T) {
		value, error := stack.Pop()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if stack.Size() != 2 {
			t.Errorf("expected size 2, got %d", stack.Size())
		}

		value, error = stack.Pop()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if stack.Size() != 1 {
			t.Errorf("expected size 1, got %d", stack.Size())
		}

		value, error = stack.Pop()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if stack.Size() != 0 {
			t.Errorf("expected size 0, got %d", stack.Size())
		}
	})

	t.Run("Pop from empty stack", func(t *testing.T) {
		_, error := stack.Pop()
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestStackPeek(t *testing.T) {
	stack := data_structures.NewStack()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Run("Peek element", func(t *testing.T) {
		value, error := stack.Peek()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if stack.Size() != 3 {
			t.Errorf("expected size 3, got %d", stack.Size())
		}
	})

	t.Run("Peek from empty stack", func(t *testing.T) {
		stack.Pop()
		stack.Pop()
		stack.Pop()
		_, error := stack.Peek()
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}
