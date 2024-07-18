package data_structures

type Color bool

const (
	Red   Color = false
	Black Color = true
)

type RBTreeNode struct {
	Value       int
	Color       Color
	Left, Right *RBTreeNode
	Parent      *RBTreeNode
}

type RBTree struct {
	Root *RBTreeNode
}

func NewRBTree() *RBTree {
	return &RBTree{}
}

func (t *RBTree) IsEmpty() bool {
	return t.Root == nil
}
