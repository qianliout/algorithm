package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	inf := 1 << 30
	var dfs func(root *TreeNode) (int, int)

	dfs = func(root *TreeNode) (int, int) {
		if root == nil {
			return -inf, 0
		}
		ly, ln := dfs(root.Left)
		ry, rn := dfs(root.Right)
		yes := root.Val + ln + rn
		no := max(ly, ln) + max(ry, rn)
		return yes, no
	}
	return max(dfs(root))
}
