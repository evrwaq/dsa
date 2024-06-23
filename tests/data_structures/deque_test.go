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
