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
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		if depth == len(ans) {
			ans = append(ans, root.Val)
		}
		dfs(root.Right, depth+1)
		dfs(root.Left, depth+1)
	}
	dfs(root, 0)
	return ans
}
