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

	tree.Insert(5)

	if tree.IsEmpty() {
		t.Error("expected tree to not be empty")
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

func TestBinaryTreeSearch(t *testing.T) {
	tree := data_structures.NewBinaryTree()

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)

	found, _ := tree.Search(5)
	if !found {
		t.Error("expected to find value 5 in the tree")
	}

	found, _ = tree.Search(3)
	if !found {
		t.Error("expected to find value 3 in the tree")
	}

	found, _ = tree.Search(7)
	if !found {
		t.Error("expected to find value 7 in the tree")
	}

	found, _ = tree.Search(10)
	if found {
		t.Error("did not expect to find value 10 in the tree")
	}
}

func TestBinaryTreeRemove(t *testing.T) {
	tree := data_structures.NewBinaryTree()

	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(7)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)
	tree.Insert(8)
	tree.Insert(2)
	tree.Insert(9)
	tree.Insert(0)

	err := tree.Remove(5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	found, _ := tree.Search(5)
	if found {
		t.Error("did not expect to find value 5 in the tree after removal")
	}

	err = tree.Remove(3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	found, _ = tree.Search(3)
	if found {
		t.Error("did not expect to find value 3 in the tree after removal")
	}

	err = tree.Remove(1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	found, _ = tree.Search(1)
	if found {
		t.Error("did not expect to find value 1 in the tree after removal")
	}

	err = tree.Remove(8)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	found, _ = tree.Search(8)
	if found {
		t.Error("did not expect to find value 8 in the tree after removal")
	}

	tree = data_structures.NewBinaryTree()
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(1)

	err = tree.Remove(3)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	found, _ = tree.Search(3)
	if found {
		t.Error("did not expect to find value 3 in the tree after removal")
	}

	err = tree.Remove(20)
	if err == nil {
		t.Error("expected error when trying to remove value not in the tree")
	}
}