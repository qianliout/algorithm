package main

import (
	"sort"
)

func main() {

}

// 会超时
func findWords(board [][]byte, words []string) []string {
	tr := Constructor()
	mx := 0
	for _, ch := range words {
		tr.Insert(ch)
		mx = max(mx, len(ch))
	}
	dirs := [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	var dfs func(i, j int)
	var word []byte
	ans := make([]string, 0)
	m, n := len(board), len(board[0])
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	add := make(map[string]bool)
	dfs = func(i, j int) {
		if i < 0 || i >= m || j < 0 || j >= n || visit[i][j] {
			return
		}
		if visit[i][j] {
			return
		}
		// 一定要加入 word，然后进行判断，再进入到下一层
		// 不然，只有一个字符的情况下，就不能获取到答案
		word = append(word, board[i][j])
		if tr.Search(string(word)) {
			if !add[string(word)] {
				ans = append(ans, string(word))
				add[string(word)] = true
			}
		}

		visit[i][j] = true
		for _, dir := range dirs {
			x, y := i+dir[0], j+dir[1]
			dfs(x, y)
		}
		word = word[:len(word)-1]
		visit[i][j] = false

	}
	for i := range board {
		for j := range board[i] {
			dfs(i, j)
		}
	}
	sort.Strings(ans)
	return ans
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

func (this *Trie) Search(word string) bool {
	var dfs func(node *Node, ss []byte, i int) bool

	dfs = func(node *Node, ss []byte, i int) bool {
		if node == nil {
			return false
		}
		if i >= len(ss) {
			return node.End
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

func gen(m, n int) [][]bool {
	visit := make([][]bool, m)
	for i := range visit {
		visit[i] = make([]bool, n)
	}
	return visit
}
