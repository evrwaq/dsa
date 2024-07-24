package data_structures

import (
	"fmt"
)

type SuffixTreeNode struct {
	Children   map[rune]*SuffixTreeNode
	Start      int
	End        *int
	SuffixLink *SuffixTreeNode
}

type SuffixTree struct {
	Root *SuffixTreeNode
	Text string
	Size int
}

func NewSuffixTree(text string) *SuffixTree {
	tree := &SuffixTree{
		Text: text,
		Size: len(text),
	}
	tree.build()
	return tree
}

func (tree *SuffixTree) build() {
	tree.Root = &SuffixTreeNode{
		Children: make(map[rune]*SuffixTreeNode),
		Start:    -1,
		End:      new(int),
	}
	*tree.Root.End = -1

	activeNode := tree.Root
	activeEdge := -1
	activeLength := 0

	remainingSuffixCount := 0
	leafEnd := -1

	for i := 0; i < tree.Size; i++ {
		previousNode := (*SuffixTreeNode)(nil)
		remainingSuffixCount++
		leafEnd = i

		for remainingSuffixCount > 0 {
			if activeLength == 0 {
				activeEdge = i
			}

			if _, ok := activeNode.Children[rune(tree.Text[activeEdge])]; !ok {
				activeNode.Children[rune(tree.Text[activeEdge])] = &SuffixTreeNode{
					Children: make(map[rune]*SuffixTreeNode),
					Start:    i,
					End:      &leafEnd,
				}

				if previousNode != nil {
					previousNode.SuffixLink = activeNode
					previousNode = nil
				}
			} else {
				next := activeNode.Children[rune(tree.Text[activeEdge])]
				edgeLength := *next.End - next.Start + 1

				if activeLength >= edgeLength {
					activeEdge += edgeLength
					activeLength -= edgeLength
					activeNode = next
					continue
				}

				if tree.Text[next.Start+activeLength] == tree.Text[i] {
					if previousNode != nil && activeNode != tree.Root {
						previousNode.SuffixLink = activeNode
						previousNode = nil
					}
					activeLength++
					break
				}

				splitEnd := next.Start + activeLength - 1
				split := &SuffixTreeNode{
					Children: make(map[rune]*SuffixTreeNode),
					Start:    next.Start,
					End:      &splitEnd,
				}
				activeNode.Children[rune(tree.Text[activeEdge])] = split
				split.Children[rune(tree.Text[i])] = &SuffixTreeNode{
					Children: make(map[rune]*SuffixTreeNode),
					Start:    i,
					End:      &leafEnd,
				}
				next.Start += activeLength
				split.Children[rune(tree.Text[next.Start])] = next

				if previousNode != nil {
					previousNode.SuffixLink = split
				}
				previousNode = split
			}

			remainingSuffixCount--

			if activeNode == tree.Root && activeLength > 0 {
				activeLength--
				activeEdge = i - remainingSuffixCount + 1
			} else if activeNode != tree.Root {
				activeNode = activeNode.SuffixLink
			}
		}
	}
}

func (tree *SuffixTree) PrintTree() {
	tree.printNode(tree.Root, 0)
}

func (tree *SuffixTree) printNode(node *SuffixTreeNode, indent int) {
	if node == nil {
		return
	}
	if node.Start != -1 {
		fmt.Printf("%*s", indent, "")
		fmt.Printf("%s\n", tree.Text[node.Start:min(*node.End+1, tree.Size)])
	}
	for _, child := range node.Children {
		tree.printNode(child, indent+2)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (tree *SuffixTree) ContainsSuffix(suffix string) bool {
	currentNode := tree.Root
	index := 0
	for index < len(suffix) {
		if node, exists := currentNode.Children[rune(suffix[index])]; exists {
			start := node.Start
			end := min(*node.End+1, tree.Size)
			if end-start > len(suffix)-index {
				end = start + len(suffix) - index
			}
			if tree.Text[start:end] != suffix[index:index+end-start] {
				return false
			}
			index += end - start
			currentNode = node
		} else {
			return false
		}
	}
	return true
}
