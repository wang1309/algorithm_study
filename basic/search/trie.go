package search

type trieNode struct {
	char interface{}
	isEnding bool
	children map[rune]*trieNode
}

func NewTrieNode(char string) *trieNode {
	return &trieNode{
		char: char,
		isEnding: false,
		children: make(map[rune]*trieNode),
	}
}

type Trie struct {
	root *trieNode
}

func NewTrie() *Trie {
	return &Trie{
		root: NewTrieNode("/"),
	}
}

// Insert 往 Trie 树中插入一个单词
func (t *Trie) Insert(word string)  {
	if len(word) == 0 {
		return
	}

	node := t.root

	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			value = NewTrieNode(string(code))
			node.children[code] = value
		}

		node = value
	}

	// 一个单词遍历完所有字符后将结尾字符打上标记
	node.isEnding = true
}

// Search 查找一个字符串
func (t *Trie) Search(word string) bool {
	if len(word) == 0 {
		return false
	}

	node := t.root

	for _, code := range word {
		value, ok := node.children[code]
		if !ok {
			return false
		}

		node = value
	}

	// 只是前缀匹配
	if node.isEnding == false {
		return false
	}

	return true
}