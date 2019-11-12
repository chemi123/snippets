package main

import (
	"fmt"
)

// 小文字のアルファベット限定のトライ木
type Trie struct {
	eow      bool
	children [26]*Trie
}

func NewTrie() Trie {
	return Trie{}
}

func (trie *Trie) Insert(word string) {
	current := trie
	for i := range word {
		if current.children[word[i]-'a'] == nil {
			current.children[word[i]-'a'] = new(Trie)
		}
		current = current.children[word[i]-'a']
	}
	current.eow = true
}

func (trie *Trie) Search(word string) bool {
	current := trie
	for i := range word {
		if current.children[word[i]-'a'] == nil {
			return false
		}
		current = current.children[word[i]-'a']
	}
	return current.eow
}

func main() {
	trie := NewTrie()
	trie.Insert("try")
	fmt.Println(trie.Search("try"))
	fmt.Println(trie.Search("trie"))
}
