package main

// https://leetcode.com/problems/design-add-and-search-words-data-structure/

type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
}

type WordDictionary struct {
	root *TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{
			children: map[rune]*TrieNode{},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
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

func (this *WordDictionary) Search(word string) bool {
	iter := []*TrieNode{this.root}

	for _, w := range word {
		if len(iter) == 0 {
			return false
		}

		nextIter := make([]*TrieNode, 0)

		for _, it := range iter {
			if w == '.' {
				for _, ch := range it.children {
					nextIter = append(nextIter, ch)
				}
				continue
			}
			if _, ok := it.children[w]; ok {
				nextIter = append(nextIter, it.children[w])
				continue
			}
		}

		iter = nextIter
	}

	for _, it := range iter {
		if it.end {
			return true
		}
	}

	return false
}

// --
// The following solution is a recursive version
// --

//func (this *WordDictionary) Search(word string) bool {
//	var dfs func(string, *TrieNode) bool
//	dfs = func(s string, root *TrieNode) bool {
//		iter := root
//
//		for i, c := range s {
//			if c == '.' {
//				for _, n := range iter.nodes {
//					if dfs(s[i+1:], n) {
//						return true
//					}
//				}
//				return false
//			} else {
//				if _, ok := iter.nodes[c]; !ok {
//					return false
//				}
//				iter = iter.nodes[c]
//			}
//		}
//
//		_, ok := iter.nodes['#']
//		return ok
//	}
//	return dfs(word, this.root)
//}
