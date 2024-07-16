package data_structures_test

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

		found, err := tree.Search(5)
		if err != nil || !found {
			t.Errorf("expected to find value 5")
		}
		found, err = tree.Search(3)
		if err != nil || !found {
			t.Errorf("expected to find value 3")
		}
		found, err = tree.Search(7)
		if err != nil || !found {
			t.Errorf("expected to find value 7")
		}
		found, err = tree.Search(1)
		if err != nil || !found {
			t.Errorf("expected to find value 1")
		}
		found, err = tree.Search(4)
		if err != nil || !found {
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
		found, err := tree.Search(5)
		if err != nil || !found {
			t.Errorf("expected to find value 5")
		}
		found, err = tree.Search(3)
		if err != nil || !found {
			t.Errorf("expected to find value 3")
		}
		found, err = tree.Search(7)
		if err != nil || !found {
			t.Errorf("expected to find value 7")
		}
	})

	t.Run("Search non-existing elements", func(t *testing.T) {
		found, err := tree.Search(10)
		if err != nil || found {
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
	tree.Insert(6)
	tree.Insert(8)

	t.Run("Remove elements", func(t *testing.T) {
		err := tree.Remove(3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		found, err := tree.Search(3)
		if err != nil || found {
			t.Errorf("expected not to find value 3")
		}

		err = tree.Remove(5)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		found, err = tree.Search(5)
		if err != nil || found {
			t.Errorf("expected not to find value 5")
		}

		err = tree.Remove(10)
		if err == nil {
			t.Error(ds_errors.ExpectedError)
		}
	})

	t.Run("Remove node with one child", func(t *testing.T) {
		tree = ds.NewBinaryTree()
		tree.Insert(5)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(1)
		tree.Insert(4)
		tree.Insert(6)
		tree.Insert(8)
		tree.Insert(2) // Ensure we have a node with one child

		err := tree.Remove(1)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		found, err := tree.Search(1)
		if err != nil || found {
			t.Errorf("expected not to find value 1")
		}
		found, err = tree.Search(2)
		if err != nil || !found {
			t.Errorf("expected to find value 2")
		}
	})

	t.Run("Remove node with two children", func(t *testing.T) {
		tree = ds.NewBinaryTree()
		tree.Insert(5)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(1)
		tree.Insert(4)
		tree.Insert(6)
		tree.Insert(8)

		err := tree.Remove(3)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		found, err := tree.Search(3)
		if err != nil || found {
			t.Errorf("expected not to find value 3")
		}
		found, err = tree.Search(4)
		if err != nil || !found {
			t.Errorf("expected to find value 4")
		}
	})

	t.Run("Find minimum", func(t *testing.T) {
		tree = ds.NewBinaryTree()
		tree.Insert(5)
		tree.Insert(3)
		tree.Insert(7)
		tree.Insert(1)
		tree.Insert(4)
		tree.Insert(6)
		tree.Insert(8)
		tree.Insert(2)

		min := ds.FindMin(tree.Root).Value
		if min != 1 {
			t.Errorf("expected min value 1, got %d", min)
		}
	})
}
