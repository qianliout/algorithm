package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minCameraCover(root *TreeNode) int {
	// byself,byfather,bychild
	var dfs func(node *TreeNode) (int, int, int)
	inf := 1 << 25

	dfs = func(node *TreeNode) (int, int, int) {
		if node == nil {
			return inf, 0, 0
		}
		l_self, l_fa, l_child := dfs(node.Left)
		r_self, r_fa, r_child := dfs(node.Right)

		byself := min(l_self, l_fa, l_child) + min(r_self, r_fa, r_child) + 1
		by_fa := min(l_self, l_child) + min(r_self, r_child)
		by_child := min(l_self+r_self, l_self+r_child, r_self+l_child)
		return byself, by_fa, by_child
	}
	i, _, i3 := dfs(root)
	return min(i, i3)
}
