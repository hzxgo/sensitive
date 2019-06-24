package trie

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// 字典树
type Trie struct {
	Root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{
		Root: NewTrieNode(0, true),
	}
}

// 批量添加短语
func (t *Trie) BatchAdd(words ...string) {
	for _, word := range words {
		t.Add(word)
	}
}

func (t *Trie) AddByLocalFile(filepath string) error {
	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, isPrefix, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			return nil
		}
		if isPrefix == true || line == nil {
			continue
		}
		t.Add(string(line))
	}
	return nil
}

// 添加单条短语
func (t *Trie) Add(word string) {
	current := t.Root
	runes := []rune(strings.TrimSpace(word))
	wordLen := len(runes)
	for i := 0; i < wordLen; i++ {
		if next, ok := current.Children[runes[i]]; ok {
			current = next
		} else {
			trieNode := NewTrieNode(runes[i], false)
			current.Children[runes[i]] = trieNode
			current = trieNode
		}
		if i == wordLen-1 {
			current.isPathEnd = true
		}
	}
}

// 校验短语是否合法，若不合法则返回false和检查到的第一个敏感词
func (t *Trie) IsValidate(word string) (bool, string) {
	var current *TrieNode
	var found bool
	var startIndex int

	runes := []rune(word)
	wordLen := len(runes)
	parent := t.Root

	for index := 0; index < wordLen; index++ {
		if current, found = parent.Children[runes[index]]; !found {
			parent = t.Root
			index = startIndex
			startIndex++
			continue
		}
		if current.IsPathEnd() && startIndex <= index {
			return false, string(runes[startIndex : index+1])
		}
		parent = current
	}

	return true, ""
}
