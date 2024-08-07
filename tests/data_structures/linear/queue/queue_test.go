package data_structures_linear_test

import (
	ds_errors "dsa/src/data_structures/errors"
	ds "dsa/src/data_structures/linear/queue"
	"testing"
)

func TestNewQueue(t *testing.T) {
	queue := ds.NewQueue()

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
	queue := ds.NewQueue()

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

func TestQueueDequeue(t *testing.T) {
	queue := ds.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	t.Run("Dequeue elements", func(t *testing.T) {
		value, error := queue.Dequeue()
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if queue.Size() != 2 {
			t.Errorf("expected size 2, got %d", queue.Size())
		}

		value, error = queue.Dequeue()
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if queue.Size() != 1 {
			t.Errorf("expected size 1, got %d", queue.Size())
		}

		value, error = queue.Dequeue()
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if queue.Size() != 0 {
			t.Errorf("expected size 0, got %d", queue.Size())
		}
	})

	t.Run("Dequeue from empty queue", func(t *testing.T) {
		_, error := queue.Dequeue()
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}

func TestQueuePeek(t *testing.T) {
	queue := ds.NewQueue()
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	t.Run("Peek element", func(t *testing.T) {
		value, error := queue.Peek()
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if queue.Size() != 3 {
			t.Errorf("expected size 3, got %d", queue.Size())
		}
	})

	t.Run("Peek from empty queue", func(t *testing.T) {
		queue.Dequeue()
		queue.Dequeue()
		queue.Dequeue()
		_, error := queue.Peek()
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}
