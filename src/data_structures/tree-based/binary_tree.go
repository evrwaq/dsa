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
