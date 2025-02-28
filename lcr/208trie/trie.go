package main

func main() {

}

type Node struct {
	End   bool
	Child []*Node
}

type Trie struct {
	Root *Node
}

func Constructor() Trie {
	return Trie{Root: &Node{}}
}

func (this *Trie) Insert(word string) {
	ss := []byte(word)
	n := len(word)
	var dfs func(o *Node, i int)
	dfs = func(o *Node, i int) {
		if o == nil {
			o = &Node{}
		}
		if i >= n {
			o.End = true
			return
		}
		idx := this.getIdx(ss[i])
		if o.Child == nil {
			o.Child = make([]*Node, 26)
		}
		if o.Child[idx] == nil {
			o.Child[idx] = &Node{}
		}
		dfs(o.Child[idx], i+1)
	}
	dfs(this.Root, 0)
}

func (this *Trie) Search(word string) bool {
	ss := []byte(word)
	n := len(word)
	var dfs func(o *Node, i int) bool
	dfs = func(o *Node, i int) bool {
		if o == nil {
			return false
		}
		if i >= n {
			return o.End
		}
		idx := this.getIdx(ss[i])
		if o.Child == nil || o.Child[idx] == nil {
			return false
		}
		return dfs(o.Child[idx], i+1)
	}

	return dfs(this.Root, 0)
}

func (this *Trie) StartsWith(prefix string) bool {
	ss := []byte(prefix)
	n := len(prefix)
	var dfs func(o *Node, i int) bool
	dfs = func(o *Node, i int) bool {
		if o == nil {
			return false
		}
		if i >= n {
			return true
		}
		idx := this.getIdx(ss[i])
		if o.Child == nil || o.Child[idx] == nil {
			return false
		}
		return dfs(o.Child[idx], i+1)
	}

	return dfs(this.Root, 0)
}

func (this *Trie) getIdx(c byte) int {
	return int(c) - int('a')
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
