package main

import (
	"fmt"
)

func main() {
	fmt.Println(minimumLengthEncoding([]string{"time", "me", "bell"}))
}

type Trie struct {
	End  bool
	Data map[byte]*Trie
}

func add(node *Trie, word []byte, idx int) {
	if idx == len(word) {
		node.End = true
		return
	}
	c := word[idx]
	if node.Data == nil {
		node.Data = make(map[byte]*Trie)
	}
	if node.Data[c] == nil {
		node.Data[c] = &Trie{}
	}
	add(node.Data[c], word, idx+1)
}

func prefix(node *Trie, word []byte, idx int) bool {
	if idx == len(word) {
		return true
	}
	c := word[idx]
	if node.Data == nil || node.Data[c] == nil {
		return false
	}
	return prefix(node.Data[c], word, idx+1)
}

func find(node *Trie, word []byte, idx int) bool {
	if idx == len(word) {
		return node.End
	}
	c := word[idx]
	if node.Data == nil || node.Data[c] == nil {
		return false
	}
	return prefix(node.Data[c], word, idx+1)
}

// 数据量不大，可以直接使用后缀 map
func minimumLengthEncoding(words []string) int {
	dic := make(map[string]bool)

	for i := range words {
		dic[words[i]] = true
	}
	for _, w := range words {
		ww := []byte(w)
		for k := 1; k <= len(ww); k++ {
			if dic[string(ww[k:])] {
				dic[string(ww[k:])] = false
			}
		}
	}
	ans := 0
	for k, v := range dic {
		if !v {
			continue
		}
		ans += len(k) + 1
	}

	return ans
}
