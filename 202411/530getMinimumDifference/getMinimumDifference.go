package main

import (
	"math"
	. "outback/algorithm/common/treenode"
)

func main() {

}

func getMinimumDifference(root *TreeNode) int {
	ans := math.MaxInt32
	var pre *TreeNode
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		if pre != nil {
			ans = min(ans, abs(node.Val-pre.Val))
		}
		pre = node
		dfs(node.Right)
	}

	dfs(root)
	return ans
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
