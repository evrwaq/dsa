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
