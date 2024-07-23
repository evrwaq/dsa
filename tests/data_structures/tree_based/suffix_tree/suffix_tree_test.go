package data_structures_tests

import (
	ds "dsa/src/data_structures/tree_based/suffix_tree"
	"testing"
)

func TestNewSuffixTree(t *testing.T) {
	st := ds.NewSuffixTree("banana")

	if st == nil {
		t.Errorf("Expected non-nil SuffixTree")
	}
}

func TestSuffixTreeStructure(t *testing.T) {
	st := ds.NewSuffixTree("banana")
	if st.root == nil {
		t.Errorf("Expected non-nil root")
	}
	if len(st.root.children) == 0 {
		t.Errorf("Expected root to have children")
	}
}

func TestSuffixTreeConstruction(t *testing.T) {
	st := ds.NewSuffixTree("banana")

	var checkNode func(node *ds.SuffixTreeNode, depth int)
	checkNode = func(node *ds.SuffixTreeNode, depth int) {
		if node == nil {
			return
		}

		for _, child := range node.children {
			if child.start == nil || child.end == nil {
				t.Errorf("Expected valid start and end for node at depth %d", depth)
			}
			checkNode(child, depth+1)
		}
	}

	checkNode(st.root, 0)
}

func TestSuffixTreeEdges(t *testing.T) {
	st := ds.NewSuffixTree("banana")

	var edgeStrings func(node *ds.SuffixTreeNode, depth int) []string
	edgeStrings = func(node *ds.SuffixTreeNode, depth int) []string {
		if node == nil {
			return []string{}
		}

		var edges []string
		for _, child := range node.children {
			if child.start != nil && child.end != nil {
				edges = append(edges, st.text[*child.start:*child.end+1])
			}
			edges = append(edges, edgeStrings(child, depth+1)...)
		}
		return edges
	}

	edges := edgeStrings(st.root, 0)

	expectedEdges := []string{"banana", "a", "na", "n", "a"}
	for _, edge := range expectedEdges {
		found := false
		for _, e := range edges {
			if e == edge {
				found = true
				break
			}
		}
		if !found {
			t.Errorf("Expected edge %s not found", edge)
		}
	}
}
