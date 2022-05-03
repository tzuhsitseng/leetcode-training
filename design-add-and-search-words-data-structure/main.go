package main

// https://leetcode.com/problems/design-add-and-search-words-data-structure/

type WordDictionary struct {
	root *TrieNode
}

type TrieNode struct {
	nodes map[rune]*TrieNode
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &TrieNode{map[rune]*TrieNode{}},
	}
}

func (this *WordDictionary) AddWord(word string) {
	iter := this.root

	for _, c := range word {
		if iter.nodes == nil {
			iter.nodes = map[rune]*TrieNode{}
		}
		if _, ok := iter.nodes[c]; !ok {
			iter.nodes[c] = &TrieNode{}
		}
		iter = iter.nodes[c]
	}

	if iter.nodes == nil {
		iter.nodes = map[rune]*TrieNode{}
	}
	if _, ok := iter.nodes['#']; !ok {
		iter.nodes['#'] = &TrieNode{}
	}
}

func (this *WordDictionary) Search(word string) bool {
	iter := []*TrieNode{this.root}

	for _, c := range word {
		candidates := make([]*TrieNode, 0)

		for _, i := range iter {
			if i.nodes == nil {
				continue
			}
			if c == '.' {
				for k, n := range i.nodes {
					if k != '#' {
						candidates = append(candidates, n)
					}
				}
				continue
			}
			if _, ok := i.nodes[c]; ok {
				candidates = append(candidates, i.nodes[c])
				continue
			}
		}

		iter = candidates
	}

	for _, i := range iter {
		if len(i.nodes) == 0 {
			continue
		}
		if _, ok := i.nodes['#']; ok {
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
