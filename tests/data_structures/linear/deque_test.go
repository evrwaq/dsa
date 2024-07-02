package data_structures_linear_test

import (
	data_structures "dsa/src/data_structures/linear"
	"testing"
)

// TestNewDeque tests the creation of a new deque and its initial size and empty state.
func TestNewDeque(t *testing.T) {
	deque := data_structures.NewDeque() // Create a new empty deque.

	// Test the initial size of the deque.
	t.Run("Initial Size", func(t *testing.T) {
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %d", deque.Size())
		}
	})

	// Test if the new deque is empty.
	t.Run("IsEmpty on new deque", func(t *testing.T) {
		if !deque.IsEmpty() {
			t.Error("expected deque to be empty")
		}
	})
}

// TestDequeAddFront tests the AddFront method of the deque.
func TestDequeAddFront(t *testing.T) {
	deque := data_structures.NewDeque() // Create a new empty deque.

	t.Run("AddFront elements", func(t *testing.T) {
		// Add elements to the front of the deque.
		deque.AddFront(1)
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		deque.AddFront(2)
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		deque.AddFront(3)
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}

		if deque.IsEmpty() {
			t.Error("expected deque to not be empty")
		}
	})
}

// TestDequeAddBack tests the AddBack method of the deque.
func TestDequeAddBack(t *testing.T) {
	deque := data_structures.NewDeque() // Create a new empty deque.

	t.Run("AddBack elements", func(t *testing.T) {
		// Add elements to the back of the deque.
		deque.AddBack(1)
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		deque.AddBack(2)
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		deque.AddBack(3)
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}

		if deque.IsEmpty() {
			t.Error("expected deque to not be empty")
		}
	})
}

// TestDequeRemoveFront tests the RemoveFront method of the deque.
func TestDequeRemoveFront(t *testing.T) {
	deque := data_structures.NewDeque() // Create a new empty deque.
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("RemoveFront elements", func(t *testing.T) {
		// Remove elements from the front of the deque and check the values and size.
		value, error := deque.RemoveFront()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		value, error = deque.RemoveFront()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		value, error = deque.RemoveFront()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %d", deque.Size())
		}
	})

	// Test removing from an empty deque.
	t.Run("RemoveFront from empty deque", func(t *testing.T) {
		_, error := deque.RemoveFront()
		if error == nil {
			t.Error(data_structures.ExpectedError)
		}
	})
}

// TestDequeRemoveBack tests the RemoveBack method of the deque.
func TestDequeRemoveBack(t *testing.T) {
	deque := data_structures.NewDeque() // Create a new empty deque.
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("RemoveBack elements", func(t *testing.T) {
		// Remove elements from the back of the deque and check the values and size.
		value, error := deque.RemoveBack()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 2 {
			t.Errorf("expected size 2, got %d", deque.Size())
		}

		value, error = deque.RemoveBack()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if deque.Size() != 1 {
			t.Errorf("expected size 1, got %d", deque.Size())
		}

		value, error = deque.RemoveBack()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 0 {
			t.Errorf("expected size 0, got %d", deque.Size())
		}
	})

	// Test removing from an empty deque.
	t.Run("RemoveBack from empty deque", func(t *testing.T) {
		_, error := deque.RemoveBack()
		if error == nil {
			t.Error(data_structures.ExpectedError)
		}
	})
}

// TestDequePeekFront tests the PeekFront method of the deque.
func TestDequePeekFront(t *testing.T) {
	deque := data_structures.NewDeque() // Create a new empty deque.
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("PeekFront element", func(t *testing.T) {
		// Peek at the front element of the deque without removing it.
		value, error := deque.PeekFront()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}
	})

	// Test peeking from an empty deque.
	t.Run("PeekFront from empty deque", func(t *testing.T) {
		deque.RemoveFront()
		deque.RemoveFront()
		deque.RemoveFront()
		_, error := deque.PeekFront()
		if error == nil {
			t.Error(data_structures.ExpectedError)
		}
	})
}

// TestDequePeekBack tests the PeekBack method of the deque.
func TestDequePeekBack(t *testing.T) {
	deque := data_structures.NewDeque() // Create a new empty deque.
	deque.AddBack(1)
	deque.AddBack(2)
	deque.AddBack(3)

	t.Run("PeekBack element", func(t *testing.T) {
		// Peek at the back element of the deque without removing it.
		value, error := deque.PeekBack()
		if error != nil {
			t.Errorf(data_structures.UnexpectedError, error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if deque.Size() != 3 {
			t.Errorf("expected size 3, got %d", deque.Size())
		}
	})

	// Test peeking from an empty deque.
	t.Run("PeekBack from empty deque", func(t *testing.T) {
		deque.RemoveBack()
		deque.RemoveBack()
		deque.RemoveBack()
		_, error := deque.PeekBack()
		if error == nil {
			t.Error(data_structures.ExpectedError)
		}
	})
}
