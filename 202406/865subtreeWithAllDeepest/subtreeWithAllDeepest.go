package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	var ans *TreeNode
	maxD := -1
	var dfs func(root *TreeNode, dep int) int

	dfs = func(root *TreeNode, dep int) int {
		if root == nil {
			maxD = max(maxD, dep)
			return dep
		}
		le := dfs(root.Left, dep+1)
		ri := dfs(root.Right, dep+1)
		if maxD == le && le == ri {
			ans = root
		}
		return max(le, ri)
	}
	dfs(root, 0)
	return ans
}
