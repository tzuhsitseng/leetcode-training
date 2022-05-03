package main

// https://leetcode.com/problems/implement-trie-prefix-tree/

type Trie struct {
	nodes map[rune]*Trie
}

func Constructor() Trie {
	return Trie{nodes: map[rune]*Trie{}}
}

func (this *Trie) Insert(word string) {
	iter := this

	for _, c := range word {
		if iter.nodes == nil {
			iter.nodes = map[rune]*Trie{}
		}
		if _, ok := iter.nodes[c]; !ok {
			iter.nodes[c] = &Trie{}
		}
		iter = iter.nodes[c]
	}

	if iter.nodes == nil {
		iter.nodes = map[rune]*Trie{}
	}
	if _, ok := iter.nodes['#']; !ok {
		iter.nodes['#'] = &Trie{}
	}
}

func (this *Trie) Search(word string) bool {
	iter := this

	for _, c := range word {
		if len(iter.nodes) == 0 {
			return false
		}
		if _, ok := iter.nodes[c]; !ok {
			return false
		}
		iter = iter.nodes[c]
	}

	_, ok := iter.nodes['#']
	return ok
}

func (this *Trie) StartsWith(prefix string) bool {
	iter := this

	for _, c := range prefix {
		if len(iter.nodes) == 0 {
			return false
		}
		if _, ok := iter.nodes[c]; !ok {
			return false
		}
		iter = iter.nodes[c]
	}

	return true
}
