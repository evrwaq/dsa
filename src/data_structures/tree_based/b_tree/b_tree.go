package data_structures

const T = 2 // Minimum degree (defines the range for number of keys)

type BTreeNode struct {
	Leaf     bool
	Keys     []int
	Children []*BTreeNode
}

type BTree struct {
	Root *BTreeNode
}

func NewBTreeNode(leaf bool) *BTreeNode {
	return &BTreeNode{
		Leaf:     leaf,
		Keys:     make([]int, 0, 2*T-1),
		Children: make([]*BTreeNode, 0, 2*T),
	}
}

func NewBTree() *BTree {
	return &BTree{
		Root: NewBTreeNode(true),
	}
}

func (t *BTree) Insert(key int) {
	root := t.Root
	if len(root.Keys) == 2*T-1 {
		s := NewBTreeNode(false)
		s.Children = append(s.Children, root)
		t.splitChild(s, 0)
		t.Root = s
		t.insertNonFull(s, key)
	} else {
		t.insertNonFull(root, key)
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
		if len(x.Children[i].Keys) == 2*T-1 {
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
	z := NewBTreeNode(y.Leaf)
	x.Children = append(x.Children[:i+1], append([]*BTreeNode{z}, x.Children[i+1:]...)...)
	x.Keys = append(x.Keys[:i], append([]int{y.Keys[T-1]}, x.Keys[i:]...)...)
	z.Keys = y.Keys[T:]
	y.Keys = y.Keys[:T-1]
	if !y.Leaf {
		z.Children = y.Children[T:]
		y.Children = y.Children[:T]
	}
}

func (t *BTree) Search(key int) (*BTreeNode, int) {
	return search(t.Root, key)
}

func search(x *BTreeNode, k int) (*BTreeNode, int) {
	i := 0
	for i < len(x.Keys) && k > x.Keys[i] {
		i++
	}
	if i < len(x.Keys) && k == x.Keys[i] {
		return x, i
	}
	if x.Leaf {
		return nil, -1
	}
	return search(x.Children[i], k)
}

func (t *BTree) Delete(key int) {
	t.delete(t.Root, key)
	if len(t.Root.Keys) == 0 && !t.Root.Leaf {
		t.Root = t.Root.Children[0]
	}
}

func (t *BTree) delete(x *BTreeNode, k int) {
	i := 0
	for i < len(x.Keys) && k > x.Keys[i] {
		i++
	}
	if i < len(x.Keys) && k == x.Keys[i] {
		if x.Leaf {
			x.Keys = append(x.Keys[:i], x.Keys[i+1:]...)
		} else {
			t.deleteInternalNode(x, k, i)
		}
	} else if !x.Leaf {
		t.deleteNonLeafNode(x, k, i)
	}
}

func (t *BTree) deleteInternalNode(x *BTreeNode, k, idx int) {
	if len(x.Children[idx].Keys) >= T {
		predecessor := t.getPredecessor(x, idx)
		x.Keys[idx] = predecessor
		t.delete(x.Children[idx], predecessor)
	} else if len(x.Children[idx+1].Keys) >= T {
		successor := t.getSuccessor(x, idx)
		x.Keys[idx] = successor
		t.delete(x.Children[idx+1], successor)
	} else {
		t.merge(x, idx)
		t.delete(x.Children[idx], k)
	}
}

func (t *BTree) deleteNonLeafNode(x *BTreeNode, k, idx int) {
	if len(x.Children[idx].Keys) < T {
		if idx > 0 && len(x.Children[idx-1].Keys) >= T {
			t.borrowFromPrev(x, idx)
		} else if idx < len(x.Keys) && len(x.Children[idx+1].Keys) >= T {
			t.borrowFromNext(x, idx)
		} else {
			if idx < len(x.Keys) {
				t.merge(x, idx)
			} else {
				t.merge(x, idx-1)
				idx--
			}
		}
	}
	t.delete(x.Children[idx], k)
}

func (t *BTree) getPredecessor(x *BTreeNode, idx int) int {
	current := x.Children[idx]
	for !current.Leaf {
		current = current.Children[len(current.Children)-1]
	}
	return current.Keys[len(current.Keys)-1]
}

func (t *BTree) getSuccessor(x *BTreeNode, idx int) int {
	current := x.Children[idx+1]
	for !current.Leaf {
		current = current.Children[0]
	}
	return current.Keys[0]
}

func (t *BTree) borrowFromPrev(x *BTreeNode, idx int) {
	child := x.Children[idx]
	sibling := x.Children[idx-1]
	child.Keys = append([]int{x.Keys[idx-1]}, child.Keys...)
	if !child.Leaf {
		child.Children = append([]*BTreeNode{sibling.Children[len(sibling.Children)-1]}, child.Children...)
	}
	x.Keys[idx-1] = sibling.Keys[len(sibling.Keys)-1]
	sibling.Keys = sibling.Keys[:len(sibling.Keys)-1]
	if !sibling.Leaf {
		sibling.Children = sibling.Children[:len(sibling.Children)-1]
	}
}

func (t *BTree) borrowFromNext(x *BTreeNode, idx int) {
	child := x.Children[idx]
	sibling := x.Children[idx+1]
	child.Keys = append(child.Keys, x.Keys[idx])
	if !child.Leaf {
		child.Children = append(child.Children, sibling.Children[0])
	}
	x.Keys[idx] = sibling.Keys[0]
	sibling.Keys = sibling.Keys[1:]
	if !sibling.Leaf {
		sibling.Children = sibling.Children[1:]
	}
}

func (t *BTree) merge(x *BTreeNode, idx int) {
	child := x.Children[idx]
	sibling := x.Children[idx+1]
	child.Keys = append(child.Keys, x.Keys[idx])
	child.Keys = append(child.Keys, sibling.Keys...)
	if !child.Leaf {
		child.Children = append(child.Children, sibling.Children...)
	}
	x.Keys = append(x.Keys[:idx], x.Keys[idx+1:]...)
	x.Children = append(x.Children[:idx+1], x.Children[idx+2:]...)
}
