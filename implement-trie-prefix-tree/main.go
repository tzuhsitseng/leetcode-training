package main

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */

type Trie struct {
	nodes map[string]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{nodes: map[string]*Trie{}}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	curr := this
	for _, v := range word {
		c := string(v)
		if curr.nodes == nil {
			curr.nodes = map[string]*Trie{}
		}
		if _, ok := curr.nodes[c]; !ok {
			curr.nodes[c] = &Trie{}
		}
		curr = curr.nodes[c]
	}
	if curr.nodes == nil {
		curr.nodes = map[string]*Trie{}
	}
	curr.nodes["#"] = &Trie{}
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	curr := this
	for _, v := range word {
		c := string(v)
		if curr.nodes == nil {
			return false
		}
		if _, ok := curr.nodes[c]; !ok {
			return false
		}
		curr = curr.nodes[c]
	}
	_, ok := curr.nodes["#"]
	return ok
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	curr := this
	for _, v := range prefix {
		c := string(v)
		if curr.nodes == nil {
			return false
		}
		if _, ok := curr.nodes[c]; !ok {
			return false
		}
		curr = curr.nodes[c]
	}
	return true
}
