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

func TestBTree_Insert(t *testing.T) {
	tree := ds.NewBTree(3)
	tree.Insert(10)
	tree.Insert(20)
	tree.Insert(5)
	tree.Insert(6)
	tree.Insert(12)
	tree.Insert(30)
	tree.Insert(7)
	tree.Insert(17)

	if len(tree.Root.Keys) != 1 {
		t.Errorf("expected 1 key in root, got %d", len(tree.Root.Keys))
	}
	if len(tree.Root.Children) != 2 {
		t.Errorf("expected 2 children in root, got %d", len(tree.Root.Children))
	}
	if tree.Root.Keys[0] != 10 {
		t.Errorf("expected root key to be 10, got %d", tree.Root.Keys[0])
	}
}
