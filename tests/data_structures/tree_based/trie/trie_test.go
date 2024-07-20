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

func TestTrieStartsWith(t *testing.T) {
	trie := ds.NewTrie()

	words := []string{"apple", "app", "application", "bat", "ball"}
	for _, word := range words {
		trie.Insert(word)
	}

	tests := []struct {
		prefix   string
		expected bool
	}{
		{"app", true},
		{"appl", true},
		{"bat", true},
		{"ba", true},
		{"cat", false},
		{"batt", false},
		{"", true},
	}

	for _, tt := range tests {
		found := trie.StartsWith(tt.prefix)
		if found != tt.expected {
			t.Errorf("expected %v, got %v for prefix %s", tt.expected, found, tt.prefix)
		}
	}
}

func TestTrieEdgeCases(t *testing.T) {
	trie := ds.NewTrie()

	if trie.Search("anything") {
		t.Error("expected false for searching any word in an empty trie")
	}

	trie.Insert("")
	if !trie.Search("") {
		t.Error("expected true for searching an empty string")
	}

	if !trie.StartsWith("") {
		t.Error("expected true for StartsWith an empty string on empty trie")
	}

	trie.Insert("a")
	if !trie.StartsWith("") {
		t.Error("expected true for StartsWith an empty string on non-empty trie")
	}
}
