package data_structures_test

import (
	ds "dsa/src/data_structures/tree_based/avl_tree"
	"testing"
)

func TestNewAVLTree(t *testing.T) {
	tree := ds.NewAVLTree()

	t.Run("Initial State", func(t *testing.T) {
		if !tree.IsEmpty() {
			t.Error("expected tree to be empty")
		}
	})
}

func TestAVLTreeInsert(t *testing.T) {
	tree := ds.NewAVLTree()

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

func TestAVLTreeSearch(t *testing.T) {
	tree := ds.NewAVLTree()
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
