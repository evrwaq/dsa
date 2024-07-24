package data_structures_tests

import (
	ds "dsa/src/data_structures/tree_based/suffix_tree"
	"testing"
)

func TestSuffixTree(t *testing.T) {
	tree := ds.NewSuffixTree("banana")

	if tree == nil {
		t.Error("Expected SuffixTree object, got nil")
	}

	if tree.Text != "banana" {
		t.Errorf("Expected text to be 'banana', got '%s'", tree.Text)
	}

	if tree.Size != 6 {
		t.Errorf("Expected size to be 6, got %d", tree.Size)
	}
}

func TestSuffixTreeStructure(t *testing.T) {
	tree := ds.NewSuffixTree("banana")

	expectedSuffixes := []string{"banana", "anana", "nana", "ana", "na", "a"}
	for _, suffix := range expectedSuffixes {
		if !tree.ContainsSuffix(suffix) {
			t.Errorf("Suffix tree does not contain expected suffix '%s'", suffix)
		}
	}
}

func TestSuffixTreePrint(t *testing.T) {
	tree := ds.NewSuffixTree("banana")
	tree.PrintTree()
}
