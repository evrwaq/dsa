package data_structures_test

import (
	"testing"

	ds "dsa/src/data_structures/tree_based/b_tree"
)

func TestBTreeInsertAndSearch(t *testing.T) {
	btree := ds.NewBTree()

	keys := []int{10, 20, 5, 6, 12, 30, 7, 17}

	for _, key := range keys {
		btree.Insert(key)
	}

	tests := []struct {
		key      int
		expected bool
	}{
		{10, true},
		{20, true},
		{5, true},
		{6, true},
		{12, true},
		{30, true},
		{7, true},
		{17, true},
		{99, false},
	}

	for _, tt := range tests {
		node, idx := btree.Search(tt.key)
		found := node != nil && node.Keys[idx] == tt.key
		if found != tt.expected {
			t.Errorf("expected %v, got %v for key %d", tt.expected, found, tt.key)
		}
	}
}

func TestBTreeDelete(t *testing.T) {
	btree := ds.NewBTree()

	keys := []int{10, 20, 5, 6, 12, 30, 7, 17, 3}

	for _, key := range keys {
		btree.Insert(key)
	}

	deleteKeys := []int{6, 13, 7, 4, 2, 16, 15, 10, 12}

	for _, key := range deleteKeys {
		btree.Delete(key)
	}

	remainingKeys := []int{3, 5, 17, 20, 30}
	for _, key := range remainingKeys {
		node, idx := btree.Search(key)
		if node == nil || node.Keys[idx] != key {
			t.Errorf("expected to find key %d", key)
		}
	}

	deletedKeys := []int{6, 13, 7, 4, 2, 16, 15, 10, 12}
	for _, key := range deletedKeys {
		node, _ := btree.Search(key)
		if node != nil {
			t.Errorf("did not expect to find key %d", key)
		}
	}
}

func TestBTreeSplit(t *testing.T) {
	btree := ds.NewBTree()

	keys := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	for _, key := range keys {
		btree.Insert(key)
	}

	for _, key := range keys {
		node, idx := btree.Search(key)
		if node == nil || node.Keys[idx] != key {
			t.Errorf("expected to find key %d", key)
		}
	}
}

func TestBTreeDeleteEdgeCases(t *testing.T) {
	btree := ds.NewBTree()

	// Edge case 1: Deleting from an empty tree
	btree.Delete(1)

	// Edge case 2: Deleting the only key in a tree
	btree.Insert(1)
	btree.Delete(1)
	if btree.Root != nil && len(btree.Root.Keys) > 0 {
		t.Error("expected an empty tree after deleting the only key")
	}

	// Edge case 3: Deleting a non-existent key
	btree.Insert(1)
	btree.Insert(2)
	btree.Delete(3) // 3 does not exist
	node, idx := btree.Search(1)
	if node == nil || node.Keys[idx] != 1 {
		t.Error("expected to find key 1")
	}
	node, idx = btree.Search(2)
	if node == nil || node.Keys[idx] != 2 {
		t.Error("expected to find key 2")
	}

	// Edge case 4: Complex deletion scenario with merges and redistributions
	btree = ds.NewBTree()
	keys := []int{10, 20, 5, 6, 12, 30, 7, 17, 3, 25, 35, 40, 50, 60}
	for _, key := range keys {
		btree.Insert(key)
	}
	var deleteKeys = []int{6, 7, 3, 5, 10, 12, 20, 25, 30, 35, 40, 50, 60}
	for _, key := range deleteKeys {
		btree.Delete(key)
	}
	var remainingKeys = []int{17}
	for _, key := range remainingKeys {
		node, idx := btree.Search(key)
		if node == nil || node.Keys[idx] != key {
			t.Errorf("expected to find key %d", key)
		}
	}

	// Ensure tree is empty after deleting all keys
	btree.Delete(17)
	if btree.Root != nil && len(btree.Root.Keys) > 0 {
		t.Error("expected an empty tree after deleting all keys")
	}
}

func TestBTreeSearchNonExistentKeys(t *testing.T) {
	btree := ds.NewBTree()
	keys := []int{1, 2, 3, 4, 5}
	for _, key := range keys {
		btree.Insert(key)
	}

	nonExistentKeys := []int{0, 6, 7, 8, 9}
	for _, key := range nonExistentKeys {
		node, _ := btree.Search(key)
		if node != nil {
			t.Errorf("did not expect to find key %d", key)
		}
	}
}

func TestBTreeEmptyTree(t *testing.T) {
	btree := ds.NewBTree()
	if btree.Root == nil {
		t.Error("expected non-nil root for a new B-Tree")
	}

	node, _ := btree.Search(1)
	if node != nil {
		t.Error("did not expect to find any key in an empty tree")
	}
}

func TestBTreeSingleNodeTree(t *testing.T) {
	btree := ds.NewBTree()
	btree.Insert(1)

	node, idx := btree.Search(1)
	if node == nil || node.Keys[idx] != 1 {
		t.Error("expected to find key 1 in the tree")
	}

	btree.Delete(1)
	node, _ = btree.Search(1)
	if node != nil {
		t.Error("did not expect to find key 1 after deletion")
	}

	if btree.Root != nil && len(btree.Root.Keys) > 0 {
		t.Error("expected an empty tree after deleting the only key")
	}
}
