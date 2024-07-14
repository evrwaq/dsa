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
	var err error
	tree.Root, err = remove(tree.Root, value)
	return err
}

func remove(node *TreeNode, value int) (*TreeNode, error) {
	if node == nil {
		return nil, errors.New(ds_errors.ValueNotFoundError)
	}
	if value < node.Value {
		var err error
		node.Left, err = remove(node.Left, value)
		return node, err
	} else if value > node.Value {
		var err error
		node.Right, err = remove(node.Right, value)
		return node, err
	} else {
		if node.Left == nil {
			return node.Right, nil
		} else if node.Right == nil {
			return node.Left, nil
		}
		minLargerNode := findMin(node.Right)
		node.Value = minLargerNode.Value
		var err error
		node.Right, err = remove(node.Right, minLargerNode.Value)
		return node, err
	}
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
