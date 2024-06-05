package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func rob(root *TreeNode) int {
	a, b := dfs(root)
	return max(a, b)
}

// 打劫root 和不打劫 root 节点的值
func dfs(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}
	lr, ln := dfs(root.Left)
	rr, rn := dfs(root.Right)
	// 打劫root 节点
	r := ln + rn + root.Val

	// 不打劫root 节点,那么两个节子点，可打劫也可以不打劫，取其中的较大值
	n := max(lr, ln) + max(rr, rn)

	return r, n
}
