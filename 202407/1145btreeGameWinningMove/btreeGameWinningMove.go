package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func btreeGameWinningMove(root *TreeNode, n int, x int) bool {
	lsz, rsz := 0, 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		ls := dfs(node.Left)
		rs := dfs(node.Right)
		if node.Val == x {
			lsz, rsz = ls, rs
		}
		return ls + rs + 1
	}
	dfs(root)
	return max(lsz, rsz, n-(lsz+rsz+1))*2 > n
}
