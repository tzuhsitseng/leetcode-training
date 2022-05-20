package main

// https://leetcode.com/problems/implement-trie-prefix-tree/

type TrieNode struct {
	end      bool
	children map[rune]*TrieNode
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			end:      false,
			children: map[rune]*TrieNode{},
		},
	}
}

func (this *Trie) Insert(word string) {
	iter := this.root

	for _, w := range word {
		if iter.children == nil {
			iter.children = map[rune]*TrieNode{}
		}
		if _, ok := iter.children[w]; !ok {
			iter.children[w] = &TrieNode{}
		}
		iter = iter.children[w]
	}

	iter.end = true
}

func (this *Trie) Search(word string) bool {
	iter := this.root

	for _, w := range word {
		if _, ok := iter.children[w]; !ok {
			return false
		}
		iter = iter.children[w]
	}

	return iter.end
}

func (this *Trie) StartsWith(prefix string) bool {
	iter := this.root

	for _, w := range prefix {
		if _, ok := iter.children[w]; !ok {
			return false
		}
		iter = iter.children[w]
	}

	return true
}
