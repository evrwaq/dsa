package bst_test

import (
	ds_errors "dsa/src/data_structures/errors"
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

func TestBinaryTreeSearch(t *testing.T) {
	tree := ds.NewBinaryTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	t.Run("Search existing elements", func(t *testing.T) {
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
	})

	t.Run("Search non-existing elements", func(t *testing.T) {
		found, error := tree.Search(10)
		if error != nil || found {
			t.Errorf("expected not to find value 10")
		}
	})
}

func TestBinaryTreeRemove(t *testing.T) {
	tree := ds.NewBinaryTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)

	t.Run("Remove elements", func(t *testing.T) {
		error := tree.Remove(3)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		found, error := tree.Search(3)
		if error != nil || found {
			t.Errorf("expected not to find value 3")
		}

		error = tree.Remove(5)
		if error != nil {
			t.Errorf("unexpected error: %v", error)
		}
		found, error = tree.Search(5)
		if error != nil || found {
			t.Errorf("expected not to find value 5")
		}

		error = tree.Remove(10)
		if error == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})
}
