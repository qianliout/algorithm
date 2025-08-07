package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func maxPathSum(root *TreeNode) int {
	ans := -1 << 32
	var dfs func(node *TreeNode) int

	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		ans = max(ans, node.Val+l, node.Val+r, node.Val+l+r, node.Val)
		return max(node.Val, node.Val+l, node.Val+r)
	}
	dfs(root)
	return ans
}
