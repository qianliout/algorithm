package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	flipMatchVoyage(root, []int{2, 1})
}

func flipMatchVoyage(root *TreeNode, voyage []int) []int {
	ans := make([]int, 0)
	var dfs func(node *TreeNode) bool

	start := 0
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return true
		}
		if node.Val != voyage[start] {
			return false
		}
		start++

		if node.Left != nil && node.Left.Val != voyage[start] {
			ans = append(ans, node.Val)
			return dfs(node.Right) && dfs(node.Left)
		}
		return dfs(node.Left) && dfs(node.Right)
	}

	can := dfs(root)
	if !can {
		return []int{-1}
	}
	return ans
}
