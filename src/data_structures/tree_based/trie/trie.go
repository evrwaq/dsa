package data_structures

type TrieNode struct {
	Children map[rune]*TrieNode
	IsEnd    bool
}

type Trie struct {
	Root *TrieNode
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		Children: make(map[rune]*TrieNode),
		IsEnd:    false,
	}
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(),
	}
}

func (t *Trie) Insert(word string) {
	node := t.Root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			node.Children[char] = NewTrieNode()
		}
		node = node.Children[char]
	}
	node.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.Root
	for _, char := range word {
		if _, exists := node.Children[char]; !exists {
			return false
		}
		node = node.Children[char]
	}
	return node.IsEnd
}
