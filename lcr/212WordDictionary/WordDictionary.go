package main

func main() {

}

type Node struct {
	End bool
	// Value byte // 边和节点值是一样的，这里存下来是了取值的时候方便一点，不存也是可以的
	Child []*Node
}

type WordDictionary struct {
	Root *Node
}

func Constructor() WordDictionary {
	return WordDictionary{Root: &Node{}}
}

func (this *WordDictionary) AddWord(word string) {
	ss := []byte(word)
	n := len(word)

	var dfs func(o *Node, idx int)

	dfs = func(o *Node, idx int) {
		if o.Child == nil {
			o.Child = make([]*Node, 26)
		}
		if idx >= n {
			o.End = true
			return
		}
		c := this.getByteInx(ss[idx])
		if o.Child[c] == nil {
			o.Child[c] = &Node{}
		}
		dfs(o.Child[c], idx+1)
	}
	dfs(this.Root, 0)
}

// 1 <= word.length <= 25
// addWord 中的 word 由小写英文字母组成
// search 中的 word 由 '.' 或小写英文字母组成
// 最多调用 104 次 addWord 和 search

func (this *WordDictionary) Search(word string) bool {
	ss := []byte(word)
	n := len(word)
	var dfs func(o *Node, idx int) bool
	dfs = func(o *Node, idx int) bool {
		if o == nil {
			return false
		}
		if idx >= n {
			return o.End
		}
		if o.Child == nil {
			return false
		}
		if ss[idx] == '.' {
			for _, no := range o.Child {
				if dfs(no, idx+1) {
					return true
				}
			}

			// 找所有
			return false
		}

		c := this.getByteInx(ss[idx])

		return dfs(o.Child[c], idx+1)
	}
	return dfs(this.Root, 0)
}

func (this *WordDictionary) getByteInx(c byte) int {
	return int(c) - int('a')
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
