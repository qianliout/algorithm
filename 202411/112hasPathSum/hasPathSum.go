package main

import (
	. "outback/algorithm/common/treenode"
)

func main() {

}

func hasPathSum(root *TreeNode, targetSum int) bool {
	var dfs func(o *TreeNode, sum int) bool
	dfs = func(o *TreeNode, sum int) bool {
		if o == nil {
			return false
		}
		if o.Left == nil && o.Right == nil {
			return o.Val == sum
		}

		if dfs(o.Left, sum-o.Val) {
			return true
		}
		return dfs(o.Right, sum-o.Val)
	}
	return dfs(root, targetSum)
}
