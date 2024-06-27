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
	Next [26]*Node
}

type Trie struct {
	Node *Node
}

type StreamChecker struct {
	trie  *Trie
	steam []byte
	Limit int
}

func (vi *Trie) add(word []byte) {
	node := vi.Node
	for i := len(word) - 1; i >= 0; i-- {
		nexId := int(word[i]) - int('a')
		if node.Next[nexId] == nil {
			node.Next[nexId] = &Node{}
		}
		node = node.Next[nexId]
	}
	node.End = true
}

func (vi *Trie) find(word []byte) bool {
	node := vi.Node
	for i := len(word) - 1; i >= 0; i-- {
		nexId := int(word[i]) - int('a')
		if node.Next[nexId] == nil {
			return false
		}
		node = node.Next[nexId]
		if node.End {
			return true
		}
	}

	return node.End
}

func Constructor(words []string) StreamChecker {
	sc := StreamChecker{
		trie:  &Trie{Node: &Node{}},
		steam: make([]byte, 0),
		Limit: 0,
	}
	for _, wo := range words {
		sc.trie.add([]byte(wo))
		sc.Limit = max(sc.Limit, len(wo))
	}
	return sc
}

func (this *StreamChecker) Query(letter byte) bool {
	this.steam = append(this.steam, letter)

	if len(this.steam) > this.Limit {
		n := len(this.steam)
		this.steam = this.steam[n-this.Limit:]
	}

	return this.trie.find(this.steam)
}
