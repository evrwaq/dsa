package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

// TestNewStack tests the creation of a new stack and its initial size and empty state.
func TestNewStack(t *testing.T) {
	stack := data_structures.NewStack() // Create a new empty stack.

	// Test the initial size of the stack.
	t.Run("Initial Size", func(t *testing.T) {
		if stack.Size() != 0 {
			t.Errorf("expected size 0, got %d", stack.Size())
		}
	})

	// Test if the new stack is empty.
	t.Run("IsEmpty on new stack", func(t *testing.T) {
		if !stack.IsEmpty() {
			t.Error("expected stack to be empty")
		}
	})
}

// TestStackPush tests the Push method of the stack.
func TestStackPush(t *testing.T) {
	stack := data_structures.NewStack() // Create a new empty stack.

	t.Run("Push elements", func(t *testing.T) {
		// Push elements onto the stack.
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

// TestStackPop tests the Pop method of the stack.
func TestStackPop(t *testing.T) {
	stack := data_structures.NewStack() // Create a new empty stack.
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Run("Pop elements", func(t *testing.T) {
		// Pop elements from the stack and check the values and size.
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

	// Test popping from an empty stack.
	t.Run("Pop from empty stack", func(t *testing.T) {
		_, error := stack.Pop()
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}

// TestStackPeek tests the Peek method of the stack.
func TestStackPeek(t *testing.T) {
	stack := data_structures.NewStack() // Create a new empty stack.
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	t.Run("Peek element", func(t *testing.T) {
		// Peek at the top element of the stack without removing it.
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

	// Test peeking from an empty stack.
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
