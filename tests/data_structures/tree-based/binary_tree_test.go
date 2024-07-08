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

func TestBinaryTreeInsert(t *testing.T) {
	tree := data_structures.NewBinaryTree()

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	if tree.Root == nil {
		t.Error("expected root to be non-nil")
	} else {
		if tree.Root.Value != 5 {
			t.Errorf("expected root value to be 5, got %v", tree.Root.Value)
		}
		if tree.Root.Left == nil || tree.Root.Left.Value != 3 {
			t.Error("expected left child value to be 3")
		}
		if tree.Root.Right == nil || tree.Root.Right.Value != 7 {
			t.Error("expected right child value to be 7")
		}
	}
}
