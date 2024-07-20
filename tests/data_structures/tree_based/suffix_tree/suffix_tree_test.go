package data_structures_test

import (
	"testing"

	ds "dsa/src/data_structures/tree_based/suffix_tree"
)

func TestSuffixTreeInsert(t *testing.T) {
	tree := ds.NewSuffixTree()
	text := "banana"

	for i := range text {
		tree.Insert(text[i:], i)
	}

	node := tree.Root
	if len(node.Children) == 0 {
		t.Error("expected non-empty root children")
	}

	for _, char := range "banana" {
		if _, exists := node.Children[char]; !exists {
			t.Errorf("expected to find character %c in the suffix tree", char)
		}
		node = node.Children[char]
	}

	if len(node.Indexes) == 0 {
		t.Error("expected indexes at the leaf node")
	}
}
