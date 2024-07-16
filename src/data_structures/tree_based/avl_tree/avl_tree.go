package data_structures

import (
	ds_errors "dsa/src/data_structures/errors"
	"errors"
)

// TreeNode represents a node in the AVL tree
type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

// AVLTree represents the AVL tree
type AVLTree struct {
	Root *TreeNode
}

// NewAVLTree creates a new AVL tree
func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// Insert inserts a value into the AVL tree
func (t *AVLTree) Insert(value int) error {
	var err error
	t.Root, err = insert(t.Root, value)
	return err
}

// insert is a recursive helper function to insert a value in the subtree rooted with node
func insert(node *TreeNode, value int) (*TreeNode, error) {
	if node == nil {
		return &TreeNode{Value: value, Height: 1}, nil
	}

	if value < node.Value {
		var err error
		node.Left, err = insert(node.Left, value)
		if err != nil {
			return nil, err
		}
	} else if value > node.Value {
		var err error
		node.Right, err = insert(node.Right, value)
		if err != nil {
			return nil, err
		}
	} else {
		return node, errors.New(ds_errors.ValueExistsError)
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	balance := getBalance(node)

	// Left Left Case
	if balance > 1 && value < node.Left.Value {
		return rightRotate(node), nil
	}

	// Right Right Case
	if balance < -1 && value > node.Right.Value {
		return leftRotate(node), nil
	}

	// Left Right Case
	if balance > 1 && value > node.Left.Value {
		node.Left = leftRotate(node.Left)
		return rightRotate(node), nil
	}

	// Right Left Case
	if balance < -1 && value < node.Right.Value {
		node.Right = rightRotate(node.Right)
		return leftRotate(node), nil
	}

	return node, nil
}

// Find finds a value in the AVL tree
func (t *AVLTree) Find(value int) (*TreeNode, bool) {
	return find(t.Root, value)
}

// find is a recursive helper function to find a value in the subtree rooted with node
func find(node *TreeNode, value int) (*TreeNode, bool) {
	if node == nil {
		return nil, false
	}

	if value < node.Value {
		return find(node.Left, value)
	} else if value > node.Value {
		return find(node.Right, value)
	} else {
		return node, true
	}
}

// Delete deletes a value from the AVL tree
func (t *AVLTree) Delete(value int) error {
	var err error
	t.Root, err = delete(t.Root, value)
	return err
}

// delete is a recursive helper function to delete a value in the subtree rooted with node
func delete(node *TreeNode, value int) (*TreeNode, error) {
	if node == nil {
		return node, nil
	}

	if value < node.Value {
		var err error
		node.Left, err = delete(node.Left, value)
		if err != nil {
			return nil, err
		}
	} else if value > node.Value {
		var err error
		node.Right, err = delete(node.Right, value)
		if err != nil {
			return nil, err
		}
	} else {
		if node.Left == nil || node.Right == nil {
			var temp *TreeNode
			if node.Left != nil {
				temp = node.Left
			} else {
				temp = node.Right
			}

			if temp == nil {
				node = nil
			} else {
				*node = *temp
			}
		} else {
			temp := minValueNode(node.Right)
			node.Value = temp.Value
			var err error
			node.Right, err = delete(node.Right, temp.Value)
			if err != nil {
				return nil, err
			}
		}
	}

	if node == nil {
		return node, nil
	}

	node.Height = 1 + max(height(node.Left), height(node.Right))
	balance := getBalance(node)

	// Left Left Case
	if balance > 1 && getBalance(node.Left) >= 0 {
		return rightRotate(node), nil
	}

	// Left Right Case
	if balance > 1 && getBalance(node.Left) < 0 {
		node.Left = leftRotate(node.Left)
		return rightRotate(node), nil
	}

	// Right Right Case
	if balance < -1 && getBalance(node.Right) <= 0 {
		return leftRotate(node), nil
	}

	// Right Left Case
	if balance < -1 && getBalance(node.Right) > 0 {
		node.Right = rightRotate(node.Right)
		return leftRotate(node), nil
	}

	return node, nil
}

// GetBalance returns the balance factor of the node
func (t *AVLTree) GetBalance(node *TreeNode) int {
	return getBalance(node)
}

// minValueNode returns the node with the minimum value found in that tree
func minValueNode(node *TreeNode) *TreeNode {
	current := node
	for current.Left != nil {
		current = current.Left
	}
	return current
}

// height returns the height of the node
func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// getBalance returns the balance factor of the node
func getBalance(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

// rightRotate performs a right rotation on the subtree rooted with y
func rightRotate(y *TreeNode) *TreeNode {
	x := y.Left
	T2 := x.Right

	x.Right = y
	y.Left = T2

	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	return x
}

// leftRotate performs a left rotation on the subtree rooted with x
func leftRotate(x *TreeNode) *TreeNode {
	y := x.Right
	T2 := y.Left

	y.Left = x
	x.Right = T2

	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	return y
}
