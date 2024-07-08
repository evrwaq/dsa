package data_structures_test

import (
	data_structures "dsa/src/data_structures/tree-based"
	"testing"
)

func TestNewBinaryTree(t *testing.T) {
	tree := data_structures.NewBinaryTree()

	if tree == nil {
		t.Error("expected tree to be non-nil")
	} else if tree.Root != nil {
		t.Error("expected root to be nil for a new tree")
	}
}

func TestBinaryTreeIsEmpty(t *testing.T) {
	tree := data_structures.NewBinaryTree()

	if !tree.IsEmpty() {
		t.Error("expected tree to be empty")
	}
}
