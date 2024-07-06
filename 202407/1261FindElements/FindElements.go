package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

type FindElements struct {
	Data map[int]bool
}

func Constructor(root *TreeNode) FindElements {

	data := map[int]bool{}

	var dfs func(node *TreeNode, pre int)

	dfs = func(node *TreeNode, pre int) {
		if node == nil {
			return
		}
		data[pre] = true
		dfs(node.Left, pre*2+1)
		dfs(node.Right, pre*2+2)

	}
	dfs(root, 0)
	return FindElements{Data: data}
}

func (this *FindElements) Find(target int) bool {
	return this.Data[target]
}
