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
