package data_structures_test

import (
	ds "dsa/src/data_structures/tree_based/b_tree"
	"testing"
)

func TestNewBTree(t *testing.T) {
	tree := ds.NewBTree(3)
	if tree.Root == nil {
		t.Error("expected root to be non-nil")
	}
	if tree.T != 3 {
		t.Errorf("expected t to be 3, got %d", tree.T)
	}
}
