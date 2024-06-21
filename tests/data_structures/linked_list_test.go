package data_structures_test

import (
	"dsa/src/data_structures"
	"testing"
)

func TestNewLinkedList(t *testing.T) {
	linked_list := data_structures.NewLinkedList()

	t.Run("Initial Size", func(t *testing.T) {
		if linked_list.Size() != 0 {
			t.Errorf("expected size 0, got %d", linked_list.Size())
		}
	})

	t.Run("IsEmpty on new list", func(t *testing.T) {
		if !linked_list.IsEmpty() {
			t.Error("expected list to be empty")
		}
	})
}
