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

func TestDequeAddBack(t *testing.T) {
	deque := data_structures.NewDeque()

	t.Run("AddBack elements", func(t *testing.T) {
		deque.AddBack(1)
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %deque", deque.Size())
		}

		deque.AddBack(2)
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %deque", deque.Size())
		}

		deque.AddBack(3)
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %deque", deque.Size())
		}

		if deque.IsEmpty() {
			t.Error("expected deque to not be empty")
		}
	})
}

func TestDequeRemoveFront(t *testing.T) {
	deque := data_structures.NewDeque()
	deque.AddFront(1)
	deque.AddFront(2)
	deque.AddFront(3)

	t.Run("RemoveFront elements", func(t *testing.T) {
		value, error := deque.RemoveFront()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %deque", deque.Size())
		}

		value, error = deque.RemoveFront()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %deque", deque.Size())
		}

		value, error = deque.RemoveFront()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %deque", deque.Size())
		}
	})

	t.Run("RemoveFront from empty deque", func(t *testing.T) {
		_, error := deque.RemoveFront()
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestDequeRemoveBack(t *testing.T) {
	deque := data_structures.NewDeque()
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("RemoveBack elements", func(t *testing.T) {
		value, error := deque.RemoveBack()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %deque", deque.Size())
		}

		value, error = deque.RemoveBack()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %deque", deque.Size())
		}

		value, error = deque.RemoveBack()
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %deque", deque.Size())
		}
	})

	t.Run("RemoveBack from empty deque", func(t *testing.T) {
		_, error := deque.RemoveBack()
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}
