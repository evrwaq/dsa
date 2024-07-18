package data_structures_test

import (
	ds "dsa/src/data_structures/tree_based/red_black_tree"
	"testing"
)

func TestRBTreeInitialEmpty(t *testing.T) {
	tree := ds.NewRBTree()
	if !tree.IsEmpty() {
		t.Error("expected the tree to be empty initially")
	}
}
