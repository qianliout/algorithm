package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func rob(root *TreeNode) int {
	var dfs func(node *TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		ly, ln := dfs(node.Left)
		ry, rn := dfs(node.Right)
		yes := node.Val + ln + rn
		no := max(ly, ln) + max(ry, rn)
		return yes, no
	}
	a, b := dfs(root)
	return max(a, b)
}
