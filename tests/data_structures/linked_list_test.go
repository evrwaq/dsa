package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	linkedList := data_structures.NewLinkedList()

	t.Run("Initial Size", func(t *testing.T) {
		if linkedList.Size() != 0 {
			t.Errorf("expected size 0, got %d", linkedList.Size())
		}
	})

	t.Run("IsEmpty on new list", func(t *testing.T) {
		if !linkedList.IsEmpty() {
			t.Error("expected list to be empty")
		}
	})
}

func TestLinkedListAdd(t *testing.T) {
	linkedList := data_structures.NewLinkedList()

	t.Run("Add elements", func(t *testing.T) {
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

func TestLinkedListGet(t *testing.T) {
	linkedList := data_structures.NewLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)

	t.Run("Get elements", func(t *testing.T) {
		value, error := linkedList.Get(0)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}

		value, error = linkedList.Get(1)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}

		value, error = linkedList.Get(2)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
	})

	t.Run("Get out of bounds", func(t *testing.T) {
		_, error := linkedList.Get(3)
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}

func TestLinkedListRemove(t *testing.T) {
	linkedList := data_structures.NewLinkedList()
	linkedList.Add(1)
	linkedList.Add(2)
	linkedList.Add(3)
	linkedList.Add(4)

	t.Run("Remove elements", func(t *testing.T) {
		value, error := linkedList.Remove(2) // Remove element from the middle
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 3 {
			t.Errorf("expected value 3, got %v", value)
		}
		if linkedList.Size() != 3 {
			t.Errorf("expected size 3, got %d", linkedList.Size())
		}

		value, error = linkedList.Remove(0)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 1 {
			t.Errorf("expected value 1, got %v", value)
		}
		if linkedList.Size() != 2 {
			t.Errorf("expected size 2, got %d", linkedList.Size())
		}

		value, error = linkedList.Remove(1)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 4 {
			t.Errorf("expected value 4, got %v", value)
		}
		if linkedList.Size() != 1 {
			t.Errorf("expected size 1, got %d", linkedList.Size())
		}

		value, error = linkedList.Remove(0)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		if value != 2 {
			t.Errorf("expected value 2, got %v", value)
		}
		if linkedList.Size() != 0 {
			t.Errorf("expected size 0, got %d", linkedList.Size())
		}
	})

	t.Run("Remove out of bounds", func(t *testing.T) {
		_, error := linkedList.Remove(0)
		if error == nil {
			t.Error("expected error, got nil")
		}
	})
}
