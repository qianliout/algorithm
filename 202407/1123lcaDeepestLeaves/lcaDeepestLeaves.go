package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (*TreeNode, int)

	dfs = func(node *TreeNode) (*TreeNode, int) {
		if node == nil {
			return node, 0
		}
		ln, ld := dfs(node.Left)
		rn, rd := dfs(node.Right)
		if ld > rd {
			return ln, ld + 1
		} else if rd > ld {
			return rn, rd + 1
		}
		return node, ld + 1
	}
	ans, _ := dfs(root)
	return ans
}
