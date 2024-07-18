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

func TestRBTreeInsert(t *testing.T) {
	tree := ds.NewRBTree()
	tree.Insert(10)
	tree.Insert(15)
	tree.Insert(5)
	tree.Insert(20)
	tree.Insert(1)
	tree.Insert(7)
	tree.Insert(12)
	tree.Insert(18)
	tree.Insert(25)

	tests := []struct {
		value    int
		expected bool
	}{
		{10, true},
		{15, true},
		{5, true},
		{20, true},
		{1, true},
		{7, true},
		{12, true},
		{18, true},
		{25, true},
		{30, false},
	}

	for _, test := range tests {
		found, _ := tree.Search(test.value)
		if found != test.expected {
			t.Errorf("expected to find value %d in the tree: %v", test.value, test.expected)
		}
	}
}
