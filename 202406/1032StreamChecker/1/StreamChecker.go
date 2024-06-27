package main

import (
	"fmt"
)

func main() {
	sc := Constructor([]string{"cd", "f", "kl"})
	fmt.Println(sc.Query('a'))
	fmt.Println(sc.Query('b'))
	fmt.Println(sc.Query('c'))
	fmt.Println(sc.Query('d'))
	fmt.Println(sc.Query('e'))
	fmt.Println(sc.Query('f'))
}

type Node struct {
	End  bool
	Next map[byte]*Node
}

type Trie struct {
	Node *Node
}

type StreamChecker struct {
	trie  *Trie
	steam []byte
	Limit int
}

func (vi *Trie) add(word string) {
	add(vi.Node, []byte(word), 0)
}

func add(node *Node, word []byte, idx int) {
	if idx == len(word) {
		node.End = true
		return
	}
	c := word[idx]
	if node.Next == nil {
		node.Next = make(map[byte]*Node)
	}
	if node.Next[c] == nil {
		node.Next[c] = &Node{}
	}

	add(node.Next[c], word, idx+1)
}

func (vi *Trie) find(word string) bool {
	return find(vi.Node, []byte(word), 0)
}

func find(node *Node, word []byte, idx int) bool {
	if node == nil {
		return false
	}
	if idx >= len(word) {
		return node.End
	}
	c := word[idx]
	if node.Next == nil || node.Next[c] == nil {
		return false
	}

	return find(node.Next[c], word, idx+1)
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{
		trie:  &Trie{Node: &Node{}},
		steam: make([]byte, 0),
		Limit: 0,
	}
	for _, wo := range words {
		sc.trie.add(wo)
		sc.Limit = max(sc.Limit, len(wo))
	}
	return sc
}

// 会超时
func (this *StreamChecker) Query(letter byte) bool {
	this.steam = append(this.steam, letter)

	if len(this.steam) > this.Limit {
		n := len(this.steam)
		this.steam = this.steam[n-this.Limit:]
	}

	n := len(this.steam)
	for i := 0; i < n; i++ {
		if this.trie.find(string(this.steam[i:])) {
			return true
		}
	}
	return false
}
