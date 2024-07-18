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

func TestRBTreeProperties(t *testing.T) {
	tree := ds.NewRBTree()
	values := []int{10, 15, 5, 20, 1, 7, 12, 18, 25}
	for _, value := range values {
		tree.Insert(value)
	}

	verifyRBTreeProperties(t, tree.Root)
}

func verifyRBTreeProperties(t *testing.T, node *ds.RBTreeNode) {
	if node == nil {
		return
	}

	if node.Color == ds.Red {
		if (node.Left != nil && node.Left.Color == ds.Red) ||
			(node.Right != nil && node.Right.Color == ds.Red) {
			t.Error("Red-Black Tree property violated: Red node has red child")
		}
	}

	verifyRBTreeProperties(t, node.Left)
	verifyRBTreeProperties(t, node.Right)
}

func TestRBTreeRotation(t *testing.T) {
	tree := ds.NewRBTree()
	values := []int{10, 20, 30}
	for _, value := range values {
		tree.Insert(value)
	}

	found, _ := tree.Search(10)
	if !found {
		t.Error("expected to find value 10 in the tree after rotation")
	}
	found, _ = tree.Search(20)
	if !found {
		t.Error("expected to find value 20 in the tree after rotation")
	}
	found, _ = tree.Search(30)
	if !found {
		t.Error("expected to find value 30 in the tree after rotation")
	}
}

func TestRBTreeComplexOperations(t *testing.T) {
	tree := ds.NewRBTree()
	values := []int{10, 85, 15, 70, 20, 60, 30, 50, 65, 80, 90, 40, 5, 55}
	for _, value := range values {
		tree.Insert(value)
	}

	for _, value := range values {
		found, _ := tree.Search(value)
		if !found {
			t.Errorf("expected to find value %d in the tree", value)
		}
	}

	removeValues := []int{10, 85, 55}
	for _, value := range removeValues {
		err := tree.Remove(value)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		found, _ := tree.Search(value)
		if found {
			t.Errorf("did not expect to find value %d in the tree after removal", value)
		}
		verifyRBTreeProperties(t, tree.Root)
	}
}
