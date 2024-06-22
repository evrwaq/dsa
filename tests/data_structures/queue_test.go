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

func TestQueueEnqueue(t *testing.T) {
	queue := data_structures.NewQueue()

	t.Run("Enqueue elements", func(t *testing.T) {
		queue.Enqueue(1)
		if queue.Size() != 1 {
			t.Errorf("expected size 1, got %d", queue.Size())
		}

		queue.Enqueue(2)
		if queue.Size() != 2 {
			t.Errorf("expected size 2, got %d", queue.Size())
		}

		queue.Enqueue(3)
		if queue.Size() != 3 {
			t.Errorf("expected size 3, got %d", queue.Size())
		}

		if queue.IsEmpty() {
			t.Error("expected queue to not be empty")
		}
	})
}
