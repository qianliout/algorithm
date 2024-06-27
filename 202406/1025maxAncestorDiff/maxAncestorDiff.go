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
func maxAncestorDiff(root *TreeNode) int {
	var ans int
	var dfs func(node *TreeNode, mi, mx int)
	dfs = func(node *TreeNode, mi, mx int) {
		if node == nil {
			return
		}
		mi = min(node.Val, mi)
		mx = max(node.Val, mx)
		ans = max(ans, node.Val-mi, mx-node.Val)
		dfs(node.Left, mi, mx)
		dfs(node.Right, mi, mx)
	}
	dfs(root, root.Val, root.Val)
	return ans
}
