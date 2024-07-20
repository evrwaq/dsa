package data_structures

type SuffixTreeNode struct {
	Children map[rune]*SuffixTreeNode
	Indexes  []int
}

type SuffixTree struct {
	Root *SuffixTreeNode
}

func NewSuffixTreeNode() *SuffixTreeNode {
	return &SuffixTreeNode{
		Children: make(map[rune]*SuffixTreeNode),
		Indexes:  []int{},
	}
}

func NewSuffixTree() *SuffixTree {
	return &SuffixTree{
		Root: NewSuffixTreeNode(),
	}
}

func (t *SuffixTree) Insert(suffix string, index int) {
	node := t.Root
	for _, char := range suffix {
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = NewSuffixTreeNode()
		}
		node = node.Children[char]
		node.Indexes = append(node.Indexes, index)
	}
}
