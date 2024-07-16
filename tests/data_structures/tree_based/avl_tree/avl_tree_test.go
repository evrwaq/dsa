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
