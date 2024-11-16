package main

import (
	"math"
	. "outback/algorithm/common/treenode"
)

func main() {

}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := math.MinInt64
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ans = max(ans, left+right+node.Val, node.Val, left+node.Val, right+node.Val)
		return max(left+node.Val, right+node.Val, node.Val)
	}
	dfs(root)

	return ans
}
