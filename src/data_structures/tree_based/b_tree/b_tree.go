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

func (t *BTree) Insert(k int) {
	root := t.Root
	if len(root.Keys) == 2*t.T-1 {
		s := &BTreeNode{
			Children: []*BTreeNode{root},
			Leaf:     false,
		}
		t.splitChild(s, 0)
		t.insertNonFull(s, k)
		t.Root = s
	} else {
		t.insertNonFull(root, k)
	}
}

func (t *BTree) insertNonFull(x *BTreeNode, k int) {
	i := len(x.Keys) - 1
	if x.Leaf {
		x.Keys = append(x.Keys, 0)
		for i >= 0 && k < x.Keys[i] {
			x.Keys[i+1] = x.Keys[i]
			i--
		}
		x.Keys[i+1] = k
	} else {
		for i >= 0 && k < x.Keys[i] {
			i--
		}
		i++
		if len(x.Children[i].Keys) == 2*t.T-1 {
			t.splitChild(x, i)
			if k > x.Keys[i] {
				i++
			}
		}
		t.insertNonFull(x.Children[i], k)
	}
}

func (t *BTree) splitChild(x *BTreeNode, i int) {
	y := x.Children[i]
	z := &BTreeNode{
		Keys:     append([]int(nil), y.Keys[t.T:]...),
		Children: append([]*BTreeNode(nil), y.Children[t.T:]...),
		Leaf:     y.Leaf,
	}
	x.Keys = append(x.Keys[:i], append([]int{y.Keys[t.T-1]}, x.Keys[i:]...)...)
	x.Children = append(x.Children[:i+1], append([]*BTreeNode{z}, x.Children[i+1:]...)...)
	y.Keys = y.Keys[:t.T-1]
	y.Children = y.Children[:t.T]
}
