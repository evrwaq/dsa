package data_structures_test

import (
	ds "dsa/src/data_structures/tree_based/red_black_tree"
	"testing"
)

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

func TestRBTreeRemove(t *testing.T) {
	tree := ds.NewRBTree()
	values := []int{10, 15, 5, 20, 1, 7, 12, 18, 25}
	for _, value := range values {
		tree.Insert(value)
	}

	removeTests := []struct {
		value int
	}{
		{10},
		{15},
		{5},
		{20},
		{1},
		{7},
		{12},
		{18},
		{25},
	}

	for _, test := range removeTests {
		err := tree.Remove(test.value)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		found, _ := tree.Search(test.value)
		if found {
			t.Errorf("did not expect to find value %d in the tree after removal", test.value)
		}
	}

	err := tree.Remove(30)
	if err == nil {
		t.Error("expected an error when trying to remove a non-existent value")
	}
}

func TestRBTreeSearch(t *testing.T) {
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

func TestRBTreeEmpty(t *testing.T) {
	tree := ds.NewRBTree()
	if !tree.IsEmpty() {
		t.Error("expected the tree to be empty")
	}

	tree.Insert(10)
	if tree.IsEmpty() {
		t.Error("expected the tree to not be empty")
	}

	tree.Remove(10)
	if !tree.IsEmpty() {
		t.Error("expected the tree to be empty after removing the only element")
	}
}
