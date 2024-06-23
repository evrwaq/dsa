package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestNewDeque(t *testing.T) {
	deque := data_structures.NewDeque()

	t.Run("Initial Size", func(t *testing.T) {
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %deque", deque.Size())
		}
	})

	t.Run("IsEmpty on new deque", func(t *testing.T) {
		if !deque.IsEmpty() {
			t.Error("expected deque to be empty")
		}
	})
}

func TestDequeAddFront(t *testing.T) {
	deque := data_structures.NewDeque()

	t.Run("AddFront elements", func(t *testing.T) {
		deque.AddFront(1)
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %deque", deque.Size())
		}

		deque.AddFront(2)
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %deque", deque.Size())
		}

		deque.AddFront(3)
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %deque", deque.Size())
		}

		if deque.IsEmpty() {
			t.Error("expected deque to not be empty")
		}
	})
}
