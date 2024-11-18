package main

func main() {

}

type WordDictionary struct {
	Trie Trie
}

func Constructor() WordDictionary {
	s := WordDictionary{Trie: ConstructorTrie()}
	return s
}

func (this *WordDictionary) AddWord(word string) {
	this.Trie.Insert(word)
}

func (this *WordDictionary) Search(word string) bool {
	return this.Trie.Search(word)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */

type Node struct {
	End   bool
	Child []*Node
}

type Trie struct {
	Root *Node
}

func ConstructorTrie() Trie {
	return Trie{Root: &Node{}}
}

func (this *Trie) Insert(word string) {
	var dfs func(node *Node, ss []byte, i int)
	dfs = func(node *Node, ss []byte, i int) {
		if i >= len(ss) {
			node.End = true
			return
		}

		if node.Child == nil {
			node.Child = make([]*Node, 26)
		}
		idx := getIdx(ss[i])
		if node.Child[idx] == nil {
			node.Child[idx] = &Node{}
		}
		c := node.Child[idx]
		dfs(c, ss, i+1)
	}
	dfs(this.Root, []byte(word), 0)
}

// 适配通配符
func (this *Trie) Search(word string) bool {
	var dfs func(node *Node, ss []byte, i int) bool

	dfs = func(node *Node, ss []byte, i int) bool {
		if node == nil {
			return false
		}
		if i >= len(ss) {
			return node.End
		}
		if ss[i] == '.' {
			ans := false
			for _, ch := range node.Child {
				ans = ans || dfs(ch, ss, i+1)
			}
			return ans
		}

		idx := getIdx(ss[i])

		if node.Child == nil || node.Child[idx] == nil {
			return false
		}
		c := node.Child[idx]
		return dfs(c, ss, i+1)
	}
	ans := dfs(this.Root, []byte(word), 0)
	return ans
}

func (this *Trie) StartsWith(prefix string) bool {
	var dfs func(node *Node, ss []byte, i int) bool

	dfs = func(node *Node, ss []byte, i int) bool {
		if node == nil {
			return false
		}
		if i >= len(ss) {
			return true
		}
		idx := getIdx(ss[i])
		if node.Child == nil || node.Child[idx] == nil {
			return false
		}
		c := node.Child[idx]
		return dfs(c, ss, i+1)
	}
	ans := dfs(this.Root, []byte(prefix), 0)
	return ans
}

func getIdx(s byte) int {
	return int(s) - int('a')
}

// 只有小写字母
