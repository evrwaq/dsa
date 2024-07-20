package data_structures_test

import (
	"testing"

	ds "dsa/src/data_structures/tree_based/trie"
)

func TestTrieInsert(t *testing.T) {
	trie := ds.NewTrie()

	words := []string{"apple", "app", "application", "bat", "ball"}

	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		node := trie.Root
		for _, char := range word {
			if _, exists := node.Children[char]; !exists {
				t.Errorf("expected to find character %c in the trie", char)
			}
			node = node.Children[char]
		}
		if !node.IsEnd {
			t.Errorf("expected end of word for %s", word)
		}
	}
}

func TestTrieSearch(t *testing.T) {
	trie := ds.NewTrie()

	words := []string{"apple", "app", "application", "bat", "ball"}
	for _, word := range words {
		trie.Insert(word)
	}

	tests := []struct {
		word     string
		expected bool
	}{
		{"apple", true},
		{"app", true},
		{"application", true},
		{"bat", true},
		{"ball", true},
		{"appl", false},
		{"batman", false},
		{"cat", false},
	}

	for _, tt := range tests {
		found := trie.Search(tt.word)
		if found != tt.expected {
			t.Errorf("expected %v, got %v for word %s", tt.expected, found, tt.word)
		}
	}
}
