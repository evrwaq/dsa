package data_structures_linear_test

import (
	ds_errors "dsa/src/data_structures/errors"
	ds "dsa/src/data_structures/linear/linked_list"
	"testing"
)

// TestNewLinkedList tests the creation of a new linked list and its initial size and empty state.
func TestNewLinkedList(t *testing.T) {
	linkedList := ds.NewLinkedList() // Create a new empty linked list.

	// Test the initial size of the linked list.
	t.Run("Initial Size", func(t *testing.T) {
		if linkedList.Size() != 0 {
			t.Errorf("expected size 0, got %d", linkedList.Size())
		}
	})

	// Test if the new linked list is empty.
	t.Run("IsEmpty on new list", func(t *testing.T) {
		if !linkedList.IsEmpty() {
			t.Error("expected list to be empty")
		}
	})
}

// TestLinkedListAdd tests the Add method of the linked list.
func TestLinkedListAdd(t *testing.T) {
	linkedList := ds.NewLinkedList() // Create a new empty linked list.

	t.Run("Add elements", func(t *testing.T) {
		// Add elements to the linked list.
		linkedList.Add(1)
		if linkedList.Size() != 1 {
			t.Errorf("expected size 1, got %d", linkedList.Size())
		}

		linkedList.Add(2)
		if linkedList.Size() != 2 {
			t.Errorf("expected size 2, got %d", linkedList.Size())
		}

		linkedList.Add(3)
		if linkedList.Size() != 3 {
			t.Errorf("expected size 3, got %d", linkedList.Size())
		}

		if linkedList.IsEmpty() {
			t.Error("expected list to not be empty")
		}
	})
}

// TestLinkedListGet tests the Get method of the linked list.
func TestLinkedListGet(t *testing.T) {
	linkedList := ds.NewLinkedList() // Create a new empty linked list.
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)

	// Test retrieving elements from the linked list.
	t.Run("Get elements", func(t *testing.T) {
		value, error := linkedList.Get(0)
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}

		value, error = linkedList.Get(1)
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}

		value, error = linkedList.Get(2)
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
	})

	// Test retrieving an element at an out-of-bounds index.
	t.Run("Get out of bounds", func(t *testing.T) {
		_, error := linkedList.Get(3)
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}

// TestLinkedListRemove tests the Remove method of the linked list.
func TestLinkedListRemove(t *testing.T) {
	linkedList := ds.NewLinkedList() // Create a new empty linked list.
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)

	// Test removing elements from the linked list.
	t.Run("Remove elements", func(t *testing.T) {
		value, error := linkedList.Remove(2) // Remove element from the middle
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if linkedList.Size() != 3 {
			t.Errorf("expected size 3, got %d", linkedList.Size())
		}

		value, error = linkedList.Remove(0) // Remove the first element
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if linkedList.Size() != 2 {
			t.Errorf("expected size 2, got %d", linkedList.Size())
		}

		value, error = linkedList.Remove(1) // Remove the last element
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 4 {
			t.Errorf("expected value 4, got %v", value)
		}
		if linkedList.Size() != 1 {
			t.Errorf("expected size 1, got %d", linkedList.Size())
		}

		value, error = linkedList.Remove(0) // Remove the only remaining element
		if error != nil {
			t.Errorf(ds_errors.UnexpectedError, error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if linkedList.Size() != 0 {
			t.Errorf("expected size 0, got %d", linkedList.Size())
		}
	})

	// Test removing an element at an out-of-bounds index.
	t.Run("Remove out of bounds", func(t *testing.T) {
		_, error := linkedList.Remove(0)
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}
