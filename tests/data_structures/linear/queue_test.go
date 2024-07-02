package data_structures_linear_test

import (
	data_structures "dsa/src/data_structures/linear"
	"testing"
)

// TestNewQueue tests the creation of a new queue and its initial size and empty state.
func TestNewQueue(t *testing.T) {
	queue := data_structures.NewQueue() // Create a new empty queue.

	// Test the initial size of the queue.
	t.Run("Initial Size", func(t *testing.T) {
		if queue.Size() != 0 {
			t.Errorf("expected size 0, got %d", queue.Size())
		}
	})

	// Test if the new queue is empty.
	t.Run("IsEmpty on new queue", func(t *testing.T) {
		if !queue.IsEmpty() {
			t.Error("expected queue to be empty")
		}
	})
}

// TestQueueEnqueue tests the Enqueue method of the queue.
func TestQueueEnqueue(t *testing.T) {
	queue := data_structures.NewQueue() // Create a new empty queue.

	t.Run("Enqueue elements", func(t *testing.T) {
		// Enqueue elements onto the queue.
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

// TestQueueDequeue tests the Dequeue method of the queue.
func TestQueueDequeue(t *testing.T) {
	queue := data_structures.NewQueue() // Create a new empty queue.
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	t.Run("Dequeue elements", func(t *testing.T) {
		// Dequeue elements from the queue and check the values and size.
		value, error := queue.Dequeue()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if queue.Size() != 2 {
			t.Errorf("expected size 2, got %d", queue.Size())
		}

		value, error = queue.Dequeue()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if queue.Size() != 1 {
			t.Errorf("expected size 1, got %d", queue.Size())
		}

		value, error = queue.Dequeue()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if queue.Size() != 0 {
			t.Errorf("expected size 0, got %d", queue.Size())
		}
	})

	// Test dequeuing from an empty queue.
	t.Run("Dequeue from empty queue", func(t *testing.T) {
		_, error := queue.Dequeue()
		if error == nil {
			t.Error(data_structures.ExpectedError)
		}
	})
}

// TestQueuePeek tests the Peek method of the queue.
func TestQueuePeek(t *testing.T) {
	queue := data_structures.NewQueue() // Create a new empty queue.
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	t.Run("Peek element", func(t *testing.T) {
		// Peek at the front element of the queue without removing it.
		value, error := queue.Peek()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if queue.Size() != 3 {
			t.Errorf("expected size 3, got %d", queue.Size())
		}
	})

	// Test peeking from an empty queue.
	t.Run("Peek from empty queue", func(t *testing.T) {
		queue.Dequeue()
		queue.Dequeue()
		queue.Dequeue()
		_, error := queue.Peek()
		if error == nil {
			t.Error(data_structures.ExpectedError)
		}
	})
}
