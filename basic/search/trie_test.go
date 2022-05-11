package search

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T)  {
	trie := NewTrie()
	words := []string{"Golang", "PHP", "Language", "Trie", "Go"}

	for _, word := range words {
		trie.Insert(word)
	}

	// 从 Trie 树中查找字符串
	term := "Gol"
	if trie.Search(term) {
		fmt.Printf("包含单词\"%s\"\n", term)
	} else {
		fmt.Printf("不包含单词\"%s\"\n", term)
	}


	term1 := "Go"
	if trie.Search(term1) {
		fmt.Printf("包含单词\"%s\"\n", term1)
	} else {
		fmt.Printf("不包含单词\"%s\"\n", term1)
	}

	term2 := "php"
	if trie.Search(term1) {
		fmt.Printf("包含单词\"%s\"\n", term2)
	} else {
		fmt.Printf("不包含单词\"%s\"\n", term2)
	}

}

