package main

import (
	"fmt"
)

func main() {
	fmt.Println(longestWord([]string{"a", "app", "appl", "ap", "apply", "apple", "bananasdfd"}))
	fmt.Println(longestWord([]string{"yo", "ew", "fc", "zrc", "yodn", "fcm", "qm", "qmo", "fcmz", "z", "ewq", "yod", "ewqz", "y"}))
}

type Tire struct {
	End  bool
	Data []*Tire
}

func add(vi *Tire, word []byte, idx int) {
	if idx >= len(word) {
		vi.End = true
		return
	}
	c := word[idx]
	if vi.Data == nil {
		vi.Data = make([]*Tire, 26)
	}
	nex := c - 'a'
	if vi.Data[nex] == nil {
		vi.Data[nex] = &Tire{}
	}
	add(vi.Data[nex], word, idx+1)
}

func longestWord(words []string) string {
	trie := &Tire{}
	for i := range words {
		add(trie, []byte(words[i]), 0)
	}
	ma := ""
	dfs(trie, &ma, []byte{})
	return ma
}

func dfs(trie *Tire, ma *string, path []byte) {
	if trie == nil {
		if len(*ma) < len(path) {
			*ma = string(path)
		}
		return
	}

	if trie.End && len(*ma) < len(path) {
		*ma = string(path)
	}
	for c, tr := range trie.Data {
		if tr == nil {
			continue
		}
		if !tr.End {
			continue
		}
		path = append(path, byte(c+'a'))
		dfs(tr, ma, path)
		path = path[:len(path)-1]
	}
}
