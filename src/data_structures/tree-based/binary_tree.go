package data_structures

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
