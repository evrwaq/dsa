package data_structures

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
