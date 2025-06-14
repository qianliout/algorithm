package main

func main() {

}

type Node struct {
	Child []*Node
	End   bool
}

type WordDictionary struct {
	Root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{Root: &Node{}}
}

func (w *WordDictionary) AddWord(word string) {
	add(w.Root, []byte(word), 0)
}

func add(node *Node, word []byte, idx int) {
	// 不可以在这里这样做
	// if node != nil {
	// 	node = &Node{}
	// }
	if idx >= len(word) {
		node.End = true
		return
	}

	c := int(word[idx]) - int('a')
	if node.Child == nil {
		node.Child = make([]*Node, 26)
	}
	if node.Child[c] == nil {
		// 在这里必须初始化好
		node.Child[c] = &Node{}
	}
	add(node.Child[c], word, idx+1)
}

func (w *WordDictionary) Search(word string) bool {
	return find(w.Root, []byte(word), 0)
}

func find(node *Node, word []byte, idx int) bool {
	if node == nil {
		return false
	}
	if idx >= len(word) {
		return node.End
	}
	if node.Child == nil {
		return false
	}
	if word[idx] != '.' {
		c := int(word[idx]) - int('a')
		no := node.Child[c]
		return find(no, word, idx+1)
	}
	ans := false
	for _, c := range node.Child {
		ans = ans || find(c, word, idx+1)
	}
	return ans
}
