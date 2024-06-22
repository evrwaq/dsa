package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := data_structures.NewQueue()

	t.Run("Initial Size", func(t *testing.T) {
		if queue.Size() != 0 {
			t.Errorf("expected size 0, got %d", queue.Size())
		}
	})

	t.Run("IsEmpty on new queue", func(t *testing.T) {
		if !queue.IsEmpty() {
			t.Error("expected queue to be empty")
		}
	})
}
