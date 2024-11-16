package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

type BSTIterator struct {
	Root  *TreeNode
	Stack []int
	Start int
	N     int
}

func Constructor(root *TreeNode) BSTIterator {
	s := BSTIterator{
		Root:  root,
		Stack: inorder(root),
		Start: 0,
	}
	s.N = len(s.Stack)
	return s
}

func (this *BSTIterator) Next() int {
	if !this.HasNext() {
		return 0
	}
	ans := this.Stack[this.Start]
	this.Start++
	return ans
}

func (this *BSTIterator) HasNext() bool {
	return this.Start < this.N
}

func inorder(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(o *TreeNode)
	dfs = func(o *TreeNode) {
		if o == nil {
			return
		}
		dfs(o.Left)
		ans = append(ans, o.Val)
		dfs(o.Right)
	}
	dfs(root)
	return ans
}
