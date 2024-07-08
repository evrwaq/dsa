package data_structures

import "errors"

type BinaryTreeNode struct {
	Value       interface{}
	Left, Right *BinaryTreeNode
}

type BinaryTree struct {
	Root *BinaryTreeNode
}

func NewBinaryTree() *BinaryTree {
	return &BinaryTree{}
}

func (tree *BinaryTree) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *BinaryTree) Insert(value interface{}) {
	tree.Root = insertBinary(tree.Root, value)
}

func insertBinary(node *BinaryTreeNode, value interface{}) *BinaryTreeNode {
	if node == nil {
		return &BinaryTreeNode{Value: value}
	}
	if value.(int) < node.Value.(int) {
		node.Left = insertBinary(node.Left, value)
	} else {
		node.Right = insertBinary(node.Right, value)
	}
	return node
}

func (tree *BinaryTree) Search(value interface{}) (bool, error) {
	return searchBinary(tree.Root, value)
}

func searchBinary(node *BinaryTreeNode, value interface{}) (bool, error) {
	if node == nil {
		return false, nil
	}
	if value.(int) == node.Value.(int) {
		return true, nil
	} else if value.(int) < node.Value.(int) {
		return searchBinary(node.Left, value)
	} else {
		return searchBinary(node.Right, value)
	}
}

func (tree *BinaryTree) Remove(value interface{}) error {
	var error error
	tree.Root, error = removeBinary(tree.Root, value)
	return error
}

func removeBinary(node *BinaryTreeNode, value interface{}) (*BinaryTreeNode, error) {
	if node == nil {
		return nil, errors.New("value not found in the tree")
	}
	if value.(int) < node.Value.(int) {
		var error error
		node.Left, error = removeBinary(node.Left, value)
		return node, error
	} else if value.(int) > node.Value.(int) {
		var error error
		node.Right, error = removeBinary(node.Right, value)
		return node, error
	} else {
		if node.Left == nil {
			return node.Right, nil
		} else if node.Right == nil {
			return node.Left, nil
		}
		minLargerNode := findMinBinary(node.Right)
		node.Value = minLargerNode.Value
		var error error
		node.Right, error = removeBinary(node.Right, minLargerNode.Value)
		return node, error
	}
}

func findMinBinary(node *BinaryTreeNode) *BinaryTreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
