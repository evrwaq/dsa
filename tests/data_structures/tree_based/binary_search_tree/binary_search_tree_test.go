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
