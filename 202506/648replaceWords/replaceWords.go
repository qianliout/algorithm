package main

import (
	"strings"
)

func main() {

}

func replaceWords(dictionary []string, sentence string) string {
	tr := &Trie{Root: &Node{}}
	for _, c := range dictionary {
		tr.Add(c)
	}
	ss := strings.Split(sentence, " ")
	ans := make([]string, 0)
	for _, ch := range ss {
		find := false
		for j := 1; j <= len(ch); j++ {
			if tr.Search(ch[:j]) {
				ans = append(ans, ch[:j])
				find = true
				break
			}
		}
		if !find {
			ans = append(ans, ch)
		}
	}
	return strings.Join(ans, " ")
}

type Node struct {
	Child []*Node
	End   bool
}

type Trie struct {
	Root *Node
}

func (t *Trie) Add(word string) {
	add(t.Root, []byte(word), 0)
}

func add(node *Node, w []byte, idx int) {
	if idx >= len(w) {
		node.End = true
		return
	}

	c := int(w[idx]) - int('a')
	if node.Child == nil {
		node.Child = make([]*Node, 26)
	}
	if node.Child[c] == nil {
		node.Child[c] = &Node{}
	}
	add(node.Child[c], w, idx+1)
}

func (t *Trie) Search(word string) bool {
	return search(t.Root, []byte(word), 0)
}

func search(node *Node, w []byte, idx int) bool {
	if node == nil {
		return false
	}
	if idx >= len(w) {
		return node.End
	}
	c := int(w[idx]) - int('a')

	if node.Child == nil || node.Child[c] == nil {
		return false
	}
	return search(node.Child[c], w, idx+1)
}
