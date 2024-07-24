// tree_test.go

package data_structures_test

import (
	ds_errors "dsa/src/data_structures/errors"
	ds "dsa/src/data_structures/tree_based/avl_tree"
	"testing"
)

func TestAVLTree_Insert(t *testing.T) {
	tree := ds.NewAVLTree()

	// Test inserting into an empty tree
	if err := tree.Insert(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tree.Root == nil {
		t.Errorf("Expected root to be not nil")
	}
	if tree.Root.Value != 10 {
		t.Errorf("Expected root value to be 10, got %d", tree.Root.Value)
	}

	// Test inserting lesser value
	if err := tree.Insert(5); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tree.Root.Left == nil {
		t.Errorf("Expected left child to be not nil")
	}
	if tree.Root.Left.Value != 5 {
		t.Errorf("Expected left child value to be 5, got %d", tree.Root.Left.Value)
	}

	// Test inserting greater value
	if err := tree.Insert(20); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tree.Root.Right == nil {
		t.Errorf("Expected right child to be not nil")
	}
	if tree.Root.Right.Value != 20 {
		t.Errorf("Expected right child value to be 20, got %d", tree.Root.Right.Value)
	}

	// Test balance property (Right Right Case)
	if err := tree.Insert(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tree.Root.Right.Right == nil {
		t.Errorf("Expected right-right child to be not nil")
	}
	if tree.Root.Right.Right.Value != 30 {
		t.Errorf("Expected right-right child value to be 30, got %d", tree.Root.Right.Right.Value)
	}

	// Test Left Left Case
	tree = ds.NewAVLTree()
	if err := tree.Insert(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(20); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tree.Root.Value != 20 {
		t.Errorf("Expected root value to be 20 after left-left rotation, got %d", tree.Root.Value)
	}

	// Test Left Right Case
	tree = ds.NewAVLTree()
	if err := tree.Insert(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(20); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tree.Root.Value != 20 {
		t.Errorf("Expected root value to be 20 after left-right rotation, got %d", tree.Root.Value)
	}

	// Test Right Left Case
	tree = ds.NewAVLTree()
	if err := tree.Insert(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(20); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if tree.Root.Value != 20 {
		t.Errorf("Expected root value to be 20 after right-left rotation, got %d", tree.Root.Value)
	}

	// Test inserting duplicate value
	if err := tree.Insert(10); err == nil || err.Error() != ds_errors.ValueExistsError {
		t.Errorf("Expected %v, got %v", ds_errors.ValueExistsError, err)
	}

	// Test duplicate insertion to trigger error propagation
	tree = ds.NewAVLTree()
	if err := tree.Insert(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(5); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(5); err == nil || err.Error() != ds_errors.ValueExistsError {
		t.Errorf("Expected %v, got %v", ds_errors.ValueExistsError, err)
	}

	// Test inserting greater value to right subtree
	tree = ds.NewAVLTree()
	if err := tree.Insert(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(20); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if err := tree.Insert(30); err == nil || err.Error() != ds_errors.ValueExistsError {
		t.Errorf("Expected %v, got %v", ds_errors.ValueExistsError, err)
	}
}

func TestAVLTree_Find(t *testing.T) {
	tree := ds.NewAVLTree()
	values := []int{10, 20, 5, 30, 15}
	for _, v := range values {
		_ = tree.Insert(v)
	}

	for _, v := range values {
		if node, found := tree.Find(v); !found || node.Value != v {
			t.Errorf("Expected to find value %d, got %v, %v", v, node, found)
		}
	}

	if _, found := tree.Find(99); found {
		t.Errorf("Expected not to find value 99")
	}
}

func TestAVLTree_Delete(t *testing.T) {
	tree := ds.NewAVLTree()
	values := []int{10, 20, 5, 30, 15}
	for _, v := range values {
		_ = tree.Insert(v)
	}

	// Test deleting leaf node
	if err := tree.Delete(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if _, found := tree.Find(10); found {
		t.Errorf("Expected not to find value 10")
	}

	// Test deleting non-existent value
	if err := tree.Delete(99); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test deleting node with one child
	if err := tree.Delete(5); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if _, found := tree.Find(5); found {
		t.Errorf("Expected not to find value 5")
	}

	// Test deleting node with two children
	if err := tree.Delete(20); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if _, found := tree.Find(20); found {
		t.Errorf("Expected not to find value 20")
	}

	// Test deleting root node
	tree = ds.NewAVLTree()
	_ = tree.Insert(10)
	if err := tree.Delete(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test Right Right Case deletion
	tree = ds.NewAVLTree()
	_ = tree.Insert(10)
	_ = tree.Insert(20)
	_ = tree.Insert(30)
	if err := tree.Delete(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test Left Left Case deletion
	tree = ds.NewAVLTree()
	_ = tree.Insert(30)
	_ = tree.Insert(20)
	_ = tree.Insert(10)
	if err := tree.Delete(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test Left Right Case deletion
	tree = ds.NewAVLTree()
	_ = tree.Insert(30)
	_ = tree.Insert(10)
	_ = tree.Insert(20)
	if err := tree.Delete(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test Right Left Case deletion
	tree = ds.NewAVLTree()
	_ = tree.Insert(10)
	_ = tree.Insert(30)
	_ = tree.Insert(20)
	if err := tree.Delete(10); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Test deletion triggering all balance cases
	tree = ds.NewAVLTree()
	nodes := []int{50, 30, 70, 20, 40, 60, 80, 10, 35, 55, 65}
	for _, v := range nodes {
		_ = tree.Insert(v)
	}
	// Trigger left-right case
	if err := tree.Delete(20); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Trigger right-left case
	if err := tree.Delete(80); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Trigger left-left case
	if err := tree.Delete(70); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	// Trigger right-right case
	if err := tree.Delete(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	// Additional test case to trigger error in left subtree deletion
	tree = ds.NewAVLTree()
	_ = tree.Insert(50)
	_ = tree.Insert(30)
	_ = tree.Insert(70)
	if err := tree.Delete(30); err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestAVLTree_Balance(t *testing.T) {
	tree := ds.NewAVLTree()
	if balance := tree.GetBalance(nil); balance != 0 {
		t.Errorf("Expected balance to be 0 for nil node, got %d", balance)
	}

	tree.Insert(10)
	if balance := tree.GetBalance(tree.Root); balance != 0 {
		t.Errorf("Expected balance to be 0 for root node, got %d", balance)
	}

	tree.Insert(20)
	if balance := tree.GetBalance(tree.Root); balance != -1 {
		t.Errorf("Expected balance to be -1 for root node after inserting 20, got %d", balance)
	}

	tree.Insert(5)
	if balance := tree.GetBalance(tree.Root); balance != 0 {
		t.Errorf("Expected balance to be 0 for root node after inserting 5, got %d", balance)
	}
}
