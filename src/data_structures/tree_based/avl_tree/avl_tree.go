package data_structures

import (
	ds_errors "dsa/src/data_structures/errors"
	"errors"
)

type TreeNode struct {
	Value  int
	Left   *TreeNode
	Right  *TreeNode
	Height int
}

type AVLTree struct {
	Root *TreeNode
}

func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

func (tree *AVLTree) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *AVLTree) Insert(value int) {
	tree.Root = insert(tree.Root, value)
}

func insert(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value, Height: 1}
	}
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else {
		node.Right = insert(node.Right, value)
	}
	return balance(node)
}

func (tree *AVLTree) Search(value int) (bool, error) {
	return search(tree.Root, value)
}

func search(node *TreeNode, value int) (bool, error) {
	if node == nil {
		return false, nil
	}
	if value == node.Value {
		return true, nil
	} else if value < node.Value {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

func (tree *AVLTree) Remove(value int) error {
	var error error
	tree.Root, error = remove(tree.Root, value)
	return error
}

func remove(node *TreeNode, value int) (*TreeNode, error) {
	if node == nil {
		return nil, errors.New(ds_errors.ValueNotFoundError)
	}
	if value < node.Value {
		var error error
		node.Left, error = remove(node.Left, value)
		if error != nil {
			return nil, error
		}
	} else if value > node.Value {
		var error error
		node.Right, error = remove(node.Right, value)
		if error != nil {
			return nil, error
		}
	} else {
		if node.Left == nil {
			return node.Right, nil
		} else if node.Right == nil {
			return node.Left, nil
		}
		minLargerNode := findMin(node.Right)
		node.Value = minLargerNode.Value
		var error error
		node.Right, error = remove(node.Right, minLargerNode.Value)
		if error != nil {
			return nil, error
		}
	}
	return balance(node), nil
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}

func balance(node *TreeNode) *TreeNode {
	updateHeight(node)
	if balanceFactor(node) > 1 {
		if balanceFactor(node.Left) < 0 {
			node.Left = rotateLeft(node.Left)
		}
		return rotateRight(node)
	}
	if balanceFactor(node) < -1 {
		if balanceFactor(node.Right) > 0 {
			node.Right = rotateRight(node.Right)
		}
		return rotateLeft(node)
	}
	return node
}

func height(node *TreeNode) int {
	if node == nil {
		return 0
	}
	return node.Height
}

func updateHeight(node *TreeNode) {
	node.Height = max(height(node.Left), height(node.Right)) + 1
}

func balanceFactor(node *TreeNode) int {
	return height(node.Left) - height(node.Right)
}

func rotateRight(y *TreeNode) *TreeNode {
	x := y.Left
	T2 := x.Right
	x.Right = y
	y.Left = T2
	updateHeight(y)
	updateHeight(x)
	return x
}

func rotateLeft(x *TreeNode) *TreeNode {
	y := x.Right
	T2 := y.Left
	y.Left = x
	x.Right = T2
	updateHeight(x)
	updateHeight(y)
	return y
}
