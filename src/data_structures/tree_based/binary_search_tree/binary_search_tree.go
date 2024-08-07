package data_structures

import (
	ds_errors "dsa/src/data_structures/errors"
	"errors"
)

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

type BinaryTree struct {
	Root *TreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *BinaryTree) Insert(value int) {
	tree.Root = insert(tree.Root, value)
}

func insert(node *TreeNode, value int) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if value < node.Value {
		node.Left = insert(node.Left, value)
	} else {
		node.Right = insert(node.Right, value)
	}
	return node
}

func (tree *BinaryTree) Search(value int) (bool, error) {
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

func (tree *BinaryTree) Remove(value int) error {
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
		return node, error
	} else if value > node.Value {
		var error error
		node.Right, error = remove(node.Right, value)
		return node, error
	} else {
		if node.Left == nil {
			return node.Right, nil
		} else if node.Right == nil {
			return node.Left, nil
		}
		minLargerNode := FindMin(node.Right)
		node.Value = minLargerNode.Value
		var error error
		node.Right, error = remove(node.Right, minLargerNode.Value)
		return node, error
	}
}

func FindMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
