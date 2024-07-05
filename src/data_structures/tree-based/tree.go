package data_structures

import "errors"

type TreeNode struct {
	Value       interface{}
	Left, Right *TreeNode
}

type Tree struct {
	Root *TreeNode
}

func NewTree() *Tree {
	return &Tree{}
}

func (tree *Tree) IsEmpty() bool {
	return tree.Root == nil
}

func (tree *Tree) Insert(value interface{}) {
	tree.Root = insert(tree.Root, value)
}

func insert(node *TreeNode, value interface{}) *TreeNode {
	if node == nil {
		return &TreeNode{Value: value}
	}
	if value.(int) < node.Value.(int) {
		node.Left = insert(node.Left, value)
	} else {
		node.Right = insert(node.Right, value)
	}
	return node
}

func (tree *Tree) Search(value interface{}) (bool, error) {
	return search(tree.Root, value)
}

func search(node *TreeNode, value interface{}) (bool, error) {
	if node == nil {
		return false, nil
	}
	if value.(int) == node.Value.(int) {
		return true, nil
	} else if value.(int) < node.Value.(int) {
		return search(node.Left, value)
	} else {
		return search(node.Right, value)
	}
}

func (tree *Tree) Remove(value interface{}) error {
	var err error
	tree.Root, err = remove(tree.Root, value)
	return err
}

func remove(node *TreeNode, value interface{}) (*TreeNode, error) {
	if node == nil {
		return nil, errors.New("value not found in the tree")
	}
	if value.(int) < node.Value.(int) {
		node.Left, _ = remove(node.Left, value)
	} else if value.(int) > node.Value.(int) {
		node.Right, _ = remove(node.Right, value)
	} else {
		if node.Left == nil {
			return node.Right, nil
		} else if node.Right == nil {
			return node.Left, nil
		}
		minLargerNode := findMin(node.Right)
		node.Value = minLargerNode.Value
		node.Right, _ = remove(node.Right, minLargerNode.Value)
	}
	return node, nil
}

func findMin(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}
	return node
}
