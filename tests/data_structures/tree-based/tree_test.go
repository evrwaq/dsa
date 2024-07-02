package data_structures_test

import (
	data_structures "dsa/src/data_structures/tree-based"
	"testing"
)

func TestNewTree(t *testing.T) {
	tree := data_structures.NewTree()

	if tree == nil {
		t.Error("expected tree to be non-nil")
	} else if tree.Root != nil {
		t.Error("expected root to be nil for a new tree")
	}
}
