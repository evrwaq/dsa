package data_structures

type BTreeNode struct {
	Keys     []int
	Children []*BTreeNode
	Leaf     bool
}

type BTree struct {
	Root *BTreeNode
	T    int
}

func NewBTree(t int) *BTree {
	return &BTree{
		Root: &BTreeNode{
			Keys:     []int{},
			Children: []*BTreeNode{},
			Leaf:     true,
		},
		T: t,
	}
}
