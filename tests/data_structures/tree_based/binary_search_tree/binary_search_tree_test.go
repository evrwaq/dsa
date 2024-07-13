package bst_test

import (
	ds "dsa/src/data_structures/tree_based/binary_search_tree"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tree := ds.NewBinaryTree()

	t.Run("Initial State", func(t *testing.T) {
		if !tree.IsEmpty() {
			t.Error("expected tree to be empty")
		}
	})
}

func TestBinaryTreeInsert(t *testing.T) {
	tree := ds.NewBinaryTree()

	t.Run("Insert elements", func(t *testing.T) {
		tree.Insert(5)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(1)
		tree.Insert(4)

		found, error := tree.Search(5)
		if error != nil || !found {
			t.Errorf("expected to find value 5")
		}
		found, error = tree.Search(3)
		if error != nil || !found {
			t.Errorf("expected to find value 3")
		}
		found, error = tree.Search(7)
		if error != nil || !found {
			t.Errorf("expected to find value 7")
		}
		found, error = tree.Search(1)
		if error != nil || !found {
			t.Errorf("expected to find value 1")
		}
		found, error = tree.Search(4)
		if error != nil || !found {
			t.Errorf("expected to find value 4")
		}
	})
}
