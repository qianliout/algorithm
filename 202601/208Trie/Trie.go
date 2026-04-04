package main

func main() {

}

type Node struct {
	End      bool
	Children []*Node
}

type Trie struct {
	Node *Node
}

func Constructor() Trie {
	return Trie{
		Node: &Node{
			Children: make([]*Node, 26),
		},
	}
}

func (this *Trie) Insert(word string) {
	insert(this.Node, []byte(word), 0)
}

func insert(node *Node, word []byte, i int) {
	if i == len(word) {
		node.End = true
		return
	}
	c := word[i]
	if node.Children == nil {
		node.Children = make([]*Node, 26)
	}
	idx := int(c - 'a')
	no := node.Children[idx]
	if no == nil {
		no = &Node{
			Children: make([]*Node, 26),
			End:      false,
		}
		node.Children[idx] = no
	}
	insert(node.Children[idx], word, i+1)
}

func (this *Trie) Search(word string) bool {
	return search(this.Node, []byte(word), 0)
}

func search(node *Node, word []byte, i int) bool {
	if node == nil {
		return false
	}
	if i == len(word) {
		return node.End
	}
	c := word[i]
	idx := int(c - 'a')
	if node.Children[idx] == nil {
		return false
	}
	return search(node.Children[idx], word, i+1)
}

func (this *Trie) StartsWith(prefix string) bool {
	return startsWith(this.Node, []byte(prefix), 0)
}

func startsWith(node *Node, prefix []byte, i int) bool {
	if node == nil {
		return false
	}
	if i == len(prefix) {
		return true
	}
	c := prefix[i]
	idx := int(c - 'a')
	if node.Children[idx] == nil {
		return false
	}
	return startsWith(node.Children[idx], prefix, i+1)
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
