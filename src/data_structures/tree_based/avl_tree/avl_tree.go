package data_structures

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
