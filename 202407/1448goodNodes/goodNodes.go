package main

import (
	"math"

	. "outback/algorithm/common/treenode"
)

func main() {

}

func goodNodes(root *TreeNode) int {
	var dfs func(node *TreeNode, preMax int)
	ans := 0

	dfs = func(node *TreeNode, preMax int) {
		if node == nil {
			return
		}
		if node.Val >= preMax {
			ans++
		}
		preMax = max(preMax, node.Val)
		dfs(node.Left, preMax)
		dfs(node.Right, preMax)
	}
	dfs(root, math.MinInt)
	return ans
}
