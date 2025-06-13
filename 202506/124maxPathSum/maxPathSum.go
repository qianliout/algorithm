package main

import (
	"math"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64
	dfs(root, &ans)
	return ans
}

func dfs(root *TreeNode, ans *int) int {
	if root == nil {
		return 0
	}
	left := dfs(root.Left, ans)
	right := dfs(root.Right, ans)

	a := max(root.Val, left+root.Val, right+root.Val, left+right+root.Val)
	*ans = max(*ans, a)

	return max(root.Val, left+root.Val, right+root.Val)
}
